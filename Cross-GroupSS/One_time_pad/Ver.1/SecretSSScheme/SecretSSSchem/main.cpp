//------------------------------------one time pad---------------------------------//
#include <iostream>
#include <fstream>
#include <vector>
#include<time.h>
#include<iostream>
#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <algorithm>

#include "ShamirSSScheme.h"
#include "file.h"
#include "key.h"
#include <string>
#include "variable.h"
#include"function.h"

using namespace std;

int main(){

	Var varia;
	varia.init();//初始化读入数据，从var.txt文件中读取预设数据，包括总share数，阈值。读入文件路径，文件输出路径等等

	string filePath = varia.SecretFilePathIn;		//要进行Secret Sharing 文件路径
	ofstream shuchu;								//输出测试时间结果
	string shuchuname = "OutPutFile\\" + varia.ResultFileOut + to_string(BitLen_128) + ".txt";
	shuchu.open(shuchuname);
	shuchu << "Read File Name:" << filePath << endl;

	vector<double> readFileTime, readShareFileTime, writeFileTime, writeShareFileTime, encryptFileTime, decryptFileTime;
	vector<double> encryptVFileTime, decryptVFileTime,genSecretHTime,genHashTime,calSabunTime1,calSabunTime2;

	UInt n, k;
	n = varia.TotalShares;
	k = varia.ThresholdShares;
	//inPutNK(n, k);		//输入n, k；n表示share的总数；k表示share的阈值
	UInt m, t;
	m = varia.TotalProviders;
	t = varia.ThresholdProviders;
	//inPutMT(m, t);		//输入m,t；m表示运营商的总数；t表示运营商的阈值
//--------------------------------------------------------------------------------------------------
	if (judgeLegal(n, k, m, t) == false){	//判断m,t；输入的n,k和m,t是否合法；如果不合法则跳出程序
		return 0;
	}
//--------------------------------------------------------------------------------------------------
	vector<UInt> x_id;// (n);
	x_id = varia.X_ID;
	//inPutX_ID(x_id);			//输入Share的ID；
	vector<UInt> y_id;// (m);
	y_id = varia.Y_ID;
	//inPutY_ID(y_id);			//输入プロバイダ的ID；
//--------------------------------------------------------------------------------------------------
//					从文件中获取素数 p
//--------------------------------------------------------------------------------------------------
	string primefile;
	cout << "the prime file is:" << to_string(PRIME_FILE_NAME) << endl;
	primefile = varia.PrimeFilePath + to_string(PRIME_FILE_NAME);
	BigNumber prime = getPrimeFromFile(primefile);	//从文件中读取素数
//--------------------------------------------------------------------------------------------------	
	for (int sum = 0; sum < varia.Count; sum++){	//循环测试 Count次，取平均值
		cout << sum << endl;
		time_t start, end;
		vector<BigNumber> AuthorityShare_Set;			//权限Share存储向量
//--------------------------------------------------------------------------------------------------
//					生成shares of vi
//--------------------------------------------------------------------------------------------------
		Key linshiKey;																//生成一个密钥|权限密码生成器
		SecByteBlock v_sec = linshiKey.genKey_Sec();								//生成SecByteBlock的v
		string v_str = linshiKey.covSecKeyToString(v_sec);							//将SecByteBlock的v转换成string类型
		BigNumber v = linshiKey.covHextoBigNumber(v_str);							//将十六进制的v转换成BigNumber类型	
		genAuShares(AuthorityShare_Set, v, t, m, y_id, prime, encryptVFileTime);			//生成vi的shares,并将生成的大素数写入prime_Au中；
		outPutAuthorityFiles(AuthorityShare_Set);									//输出Authority到硬盘上
//--------------------------------------------------------------------------------------------------
//					生成one time pad r
//--------------------------------------------------------------------------------------------------
		start = clock();
		SecByteBlock r_sec = linshiKey.genKey_Sec();						//生成one time pad r;
		string r_str = linshiKey.covSecKeyToString(r_sec);
		BigNumber r = linshiKey.covHextoBigNumber(r_str);
		end = clock();
		outPutKey(r, varia);												//输出密钥到硬盘上
		outPutTime(start, end, "生成One time pad r ", genSecretHTime);	
//--------------------------------------------------------------------------------------------------
//			生成q = v xor r;  进行异或操作
//--------------------------------------------------------------------------------------------------
		start = clock();
		BigNumber q = linshiKey.genOneTimePad(v, r);						//计算q = v + r 的型
		end = clock();
		outPutTime(start, end, "生成One time pad q ", genHashTime);
		v_str.~string();
//--------------------------------------------------------------------------------------------------
//			从硬盘写入内存
//--------------------------------------------------------------------------------------------------
		start = clock();
		File bigFile;
		readFile(bigFile,filePath);													//将文件从硬盘中读到内存中
		end = clock();
		outPutTime(start, end, "普通文件读入时间", readFileTime);
//-------------------------------------------------------------------------
//					p=s-q;对于word_set中每个s进行 p=s-q;运算
//-------------------------------------------------------------------------
		start = clock();
		//sMinusQ(bigFile, q);
		pPlusQ(bigFile, q);															// s=p+q;对于每个s=p+q;运算；
		end = clock();
		outPutTime(start, end, "自己提案进行差分运算时间", calSabunTime1);
//------------------------------------------------------------------------
//					Shamir's  Secret Sharing Scheme	of p
//------------------------------------------------------------------------
		vector<File> FileSet(n);													//生成n份存储share的文件
		start = clock();
		secretSharing(FileSet, bigFile, n, k, x_id, prime);							//进行secretsharing，生成的share文件存放到FileSet中；
		end = clock();
		outPutTime(start, end, "普通文件加密时间", encryptFileTime);
		start = clock();
		outPutshareFiles(FileSet);													//将shareFiles写到硬盘上
		end = clock();
		outPutTime(start, end, "输出share文件时间", writeShareFileTime);
		bigFile.clear();
		FileSet.~vector<File>();//清空
//------------------------------------------------------------------------
//					Shamir's secret reconstruction of pi
//------------------------------------------------------------------------
		File bigOutFile;																//输出复元大文件
		vector<File> shareInFile(k);
		start = clock();
		readShareFiles(shareInFile);													//读入share文件
		end = clock();
		outPutTime(start, end, "读入share文件时间", readShareFileTime);
		start = clock();
		bigOutFile = fileShamirSSS_Reconstruction(shareInFile, n, k, x_id, prime);		//进行从share文件复元明文；
		end = clock();
		outPutTime(start, end, "普通文件解密时间", decryptFileTime);
		shareInFile.~vector<File>();// 清空
//----------------------------------------------------------------------------------------
//					从vi恢复v;并使用 q = v xor r 进行计算；
//----------------------------------------------------------------------------------------
		ShamirSSScheme ReAuthorityTool;
		vector<BigNumber> FileAuthShare_Set;
		readAuthrotiy(FileAuthShare_Set,varia);
		start = clock();
		secretReconstruction_V(ReAuthorityTool, m, t, prime, FileAuthShare_Set, y_id);	//对权限秘密v进行复元；
		end = clock();
		AuthorityShare_Set.~vector<BigNumber>();
		FileAuthShare_Set.~vector<BigNumber>();
		outPutTime(start, end, "对v复元时间", decryptVFileTime);
		BigNumber v_re = ReAuthorityTool.get_secret();							//获取复元之后得到的v；
		string v_re_str = linshiKey.covBigNumberToString(v_re);					//对v数据类型进行转换；
		BigNumber re_r = readKey(varia);
		BigNumber q_re = linshiKey.genOneTimePad(v_re, re_r);					//将复元之后得到的v进行q=fh(v)运算
//----------------------------------------------------------------------------------------		
		start = clock();
		//pPlusQ(bigOutFile, q_re);
		sMinusQ(bigOutFile, q_re);
		end = clock();
		outPutTime(start, end, "自己提案进行差分运算时间", calSabunTime2);
		start = clock();
		outPutFile(bigOutFile,varia.SecretFilePathOut);													//将文件写到硬盘上;
		end = clock();
		outPutTime(start, end, "写到硬盘上时间", writeFileTime);		
	}
//----------------------------------------------------------------------------------------
//			输出统计结果
//----------------------------------------------------------------------------------------
	shuchu << BitLen_136 - 8 <<"bit"<< endl;
	outputResult(shuchu, "读入普通文件平均时间:", readFileTime,varia);
	outputResult(shuchu, "读入Share文件平均时间:", readShareFileTime, varia);
	outputResult(shuchu, "输出普通文件平均时间:", writeFileTime, varia);
	outputResult(shuchu, "输出输出Share文件平均时间:", writeShareFileTime, varia);
	outputResult(shuchu, "加密文件平均时间:", encryptFileTime, varia);
	outputResult(shuchu, "解密文件平均时间:", decryptFileTime, varia);
	outputResult(shuchu, "生成Share V平均时间:", encryptVFileTime, varia);
	outputResult(shuchu, "重构Share V平均时间:", decryptVFileTime, varia);
	outputResult(shuchu, "生成One time p 平均时间:", genSecretHTime, varia);
	outputResult(shuchu, "生成One time pad q:", genHashTime, varia);
	outputResult(shuchu, "生成差分1平均时间:", calSabunTime1, varia);
	outputResult(shuchu, "生成差分2平均时间:", calSabunTime2, varia);

	shuchu.close();

	return 0;
}

