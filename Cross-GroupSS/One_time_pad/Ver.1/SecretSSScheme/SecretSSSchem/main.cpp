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
	varia.init();//��ʼ���������ݣ���var.txt�ļ��ж�ȡԤ�����ݣ�������share������ֵ�������ļ�·�����ļ����·���ȵ�

	string filePath = varia.SecretFilePathIn;		//Ҫ����Secret Sharing �ļ�·��
	ofstream shuchu;								//�������ʱ����
	string shuchuname = "OutPutFile\\" + varia.ResultFileOut + to_string(BitLen_128) + ".txt";
	shuchu.open(shuchuname);
	shuchu << "Read File Name:" << filePath << endl;

	vector<double> readFileTime, readShareFileTime, writeFileTime, writeShareFileTime, encryptFileTime, decryptFileTime;
	vector<double> encryptVFileTime, decryptVFileTime,genSecretHTime,genHashTime,calSabunTime1,calSabunTime2;

	UInt n, k;
	n = varia.TotalShares;
	k = varia.ThresholdShares;
	//inPutNK(n, k);		//����n, k��n��ʾshare��������k��ʾshare����ֵ
	UInt m, t;
	m = varia.TotalProviders;
	t = varia.ThresholdProviders;
	//inPutMT(m, t);		//����m,t��m��ʾ��Ӫ�̵�������t��ʾ��Ӫ�̵���ֵ
//--------------------------------------------------------------------------------------------------
	if (judgeLegal(n, k, m, t) == false){	//�ж�m,t�������n,k��m,t�Ƿ�Ϸ���������Ϸ�����������
		return 0;
	}
//--------------------------------------------------------------------------------------------------
	vector<UInt> x_id;// (n);
	x_id = varia.X_ID;
	//inPutX_ID(x_id);			//����Share��ID��
	vector<UInt> y_id;// (m);
	y_id = varia.Y_ID;
	//inPutY_ID(y_id);			//����ץ�Х�����ID��
//--------------------------------------------------------------------------------------------------
//					���ļ��л�ȡ���� p
//--------------------------------------------------------------------------------------------------
	string primefile;
	cout << "the prime file is:" << to_string(PRIME_FILE_NAME) << endl;
	primefile = varia.PrimeFilePath + to_string(PRIME_FILE_NAME);
	BigNumber prime = getPrimeFromFile(primefile);	//���ļ��ж�ȡ����
//--------------------------------------------------------------------------------------------------	
	for (int sum = 0; sum < varia.Count; sum++){	//ѭ������ Count�Σ�ȡƽ��ֵ
		cout << sum << endl;
		time_t start, end;
		vector<BigNumber> AuthorityShare_Set;			//Ȩ��Share�洢����
//--------------------------------------------------------------------------------------------------
//					����shares of vi
//--------------------------------------------------------------------------------------------------
		Key linshiKey;																//����һ����Կ|Ȩ������������
		SecByteBlock v_sec = linshiKey.genKey_Sec();								//����SecByteBlock��v
		string v_str = linshiKey.covSecKeyToString(v_sec);							//��SecByteBlock��vת����string����
		BigNumber v = linshiKey.covHextoBigNumber(v_str);							//��ʮ�����Ƶ�vת����BigNumber����	
		genAuShares(AuthorityShare_Set, v, t, m, y_id, prime, encryptVFileTime);			//����vi��shares,�������ɵĴ�����д��prime_Au�У�
		outPutAuthorityFiles(AuthorityShare_Set);									//���Authority��Ӳ����
//--------------------------------------------------------------------------------------------------
//					����one time pad r
//--------------------------------------------------------------------------------------------------
		start = clock();
		SecByteBlock r_sec = linshiKey.genKey_Sec();						//����one time pad r;
		string r_str = linshiKey.covSecKeyToString(r_sec);
		BigNumber r = linshiKey.covHextoBigNumber(r_str);
		end = clock();
		outPutKey(r, varia);												//�����Կ��Ӳ����
		outPutTime(start, end, "����One time pad r ", genSecretHTime);	
//--------------------------------------------------------------------------------------------------
//			����q = v xor r;  ����������
//--------------------------------------------------------------------------------------------------
		start = clock();
		BigNumber q = linshiKey.genOneTimePad(v, r);						//����q = v + r ����
		end = clock();
		outPutTime(start, end, "����One time pad q ", genHashTime);
		v_str.~string();
//--------------------------------------------------------------------------------------------------
//			��Ӳ��д���ڴ�
//--------------------------------------------------------------------------------------------------
		start = clock();
		File bigFile;
		readFile(bigFile,filePath);													//���ļ���Ӳ���ж����ڴ���
		end = clock();
		outPutTime(start, end, "��ͨ�ļ�����ʱ��", readFileTime);
//-------------------------------------------------------------------------
//					p=s-q;����word_set��ÿ��s���� p=s-q;����
//-------------------------------------------------------------------------
		start = clock();
		//sMinusQ(bigFile, q);
		pPlusQ(bigFile, q);															// s=p+q;����ÿ��s=p+q;���㣻
		end = clock();
		outPutTime(start, end, "�Լ��᰸���в������ʱ��", calSabunTime1);
//------------------------------------------------------------------------
//					Shamir's  Secret Sharing Scheme	of p
//------------------------------------------------------------------------
		vector<File> FileSet(n);													//����n�ݴ洢share���ļ�
		start = clock();
		secretSharing(FileSet, bigFile, n, k, x_id, prime);							//����secretsharing�����ɵ�share�ļ���ŵ�FileSet�У�
		end = clock();
		outPutTime(start, end, "��ͨ�ļ�����ʱ��", encryptFileTime);
		start = clock();
		outPutshareFiles(FileSet);													//��shareFilesд��Ӳ����
		end = clock();
		outPutTime(start, end, "���share�ļ�ʱ��", writeShareFileTime);
		bigFile.clear();
		FileSet.~vector<File>();//���
//------------------------------------------------------------------------
//					Shamir's secret reconstruction of pi
//------------------------------------------------------------------------
		File bigOutFile;																//�����Ԫ���ļ�
		vector<File> shareInFile(k);
		start = clock();
		readShareFiles(shareInFile);													//����share�ļ�
		end = clock();
		outPutTime(start, end, "����share�ļ�ʱ��", readShareFileTime);
		start = clock();
		bigOutFile = fileShamirSSS_Reconstruction(shareInFile, n, k, x_id, prime);		//���д�share�ļ���Ԫ���ģ�
		end = clock();
		outPutTime(start, end, "��ͨ�ļ�����ʱ��", decryptFileTime);
		shareInFile.~vector<File>();// ���
//----------------------------------------------------------------------------------------
//					��vi�ָ�v;��ʹ�� q = v xor r ���м��㣻
//----------------------------------------------------------------------------------------
		ShamirSSScheme ReAuthorityTool;
		vector<BigNumber> FileAuthShare_Set;
		readAuthrotiy(FileAuthShare_Set,varia);
		start = clock();
		secretReconstruction_V(ReAuthorityTool, m, t, prime, FileAuthShare_Set, y_id);	//��Ȩ������v���и�Ԫ��
		end = clock();
		AuthorityShare_Set.~vector<BigNumber>();
		FileAuthShare_Set.~vector<BigNumber>();
		outPutTime(start, end, "��v��Ԫʱ��", decryptVFileTime);
		BigNumber v_re = ReAuthorityTool.get_secret();							//��ȡ��Ԫ֮��õ���v��
		string v_re_str = linshiKey.covBigNumberToString(v_re);					//��v�������ͽ���ת����
		BigNumber re_r = readKey(varia);
		BigNumber q_re = linshiKey.genOneTimePad(v_re, re_r);					//����Ԫ֮��õ���v����q=fh(v)����
//----------------------------------------------------------------------------------------		
		start = clock();
		//pPlusQ(bigOutFile, q_re);
		sMinusQ(bigOutFile, q_re);
		end = clock();
		outPutTime(start, end, "�Լ��᰸���в������ʱ��", calSabunTime2);
		start = clock();
		outPutFile(bigOutFile,varia.SecretFilePathOut);													//���ļ�д��Ӳ����;
		end = clock();
		outPutTime(start, end, "д��Ӳ����ʱ��", writeFileTime);		
	}
//----------------------------------------------------------------------------------------
//			���ͳ�ƽ��
//----------------------------------------------------------------------------------------
	shuchu << BitLen_136 - 8 <<"bit"<< endl;
	outputResult(shuchu, "������ͨ�ļ�ƽ��ʱ��:", readFileTime,varia);
	outputResult(shuchu, "����Share�ļ�ƽ��ʱ��:", readShareFileTime, varia);
	outputResult(shuchu, "�����ͨ�ļ�ƽ��ʱ��:", writeFileTime, varia);
	outputResult(shuchu, "������Share�ļ�ƽ��ʱ��:", writeShareFileTime, varia);
	outputResult(shuchu, "�����ļ�ƽ��ʱ��:", encryptFileTime, varia);
	outputResult(shuchu, "�����ļ�ƽ��ʱ��:", decryptFileTime, varia);
	outputResult(shuchu, "����Share Vƽ��ʱ��:", encryptVFileTime, varia);
	outputResult(shuchu, "�ع�Share Vƽ��ʱ��:", decryptVFileTime, varia);
	outputResult(shuchu, "����One time p ƽ��ʱ��:", genSecretHTime, varia);
	outputResult(shuchu, "����One time pad q:", genHashTime, varia);
	outputResult(shuchu, "���ɲ��1ƽ��ʱ��:", calSabunTime1, varia);
	outputResult(shuchu, "���ɲ��2ƽ��ʱ��:", calSabunTime2, varia);

	shuchu.close();

	return 0;
}

