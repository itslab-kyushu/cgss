#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<vector>
#include "variable.h"

using namespace std;

class ShamirSSScheme{
	
	public:

		void init(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET);		//初始化程序;	
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, BigNrVec& sharePart);	//配置参数1
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime);	//配置参数2
		void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, bool whetherGenPoly);//配置参数3
		void write_m_nK(UInt nK);						//写阈值
		void write_m_nN(UInt nN);						//写share总数值
		void write_m_bnPrimeNr(BigNumber m_bnPrime);	//写大素数值
		void write_x_id_set(vector<UInt>& x_id_set);		//写x_id向量值 
		void write_x_id_set_i(UInt i, UInt data);		//写x_id的i号位置值
		void write_m_vSharingPart(BigNrVec& sharePart);	//写入sharePart
		void write_secret(BigNumber s);							//写入secret;


		const BigNumber get_m_bnPrimeNr();		//取大素数值
		const vector<UInt>& get_x_id_set();		//取x_id向量
		const UInt get_x_id_set_i(UInt i);		//取x_id的i号位置值
		const BigNrVec& get_m_vSharing();		//取sharePart;
		const BigNumber get_secret();					//取secret;

		void gen_Prime();				//生成大素数
		void gen_Polynom();				//生成多项式
		void gen_SecretSharePart();		//生成share
		void secret_Reconstruction();   //重新构造秘密

	private:
		UInt m_nN; //share总数
		UInt m_nK; //阈值
		BigNumber m_bnPrimeNr; //大素数
		BigNrVec m_vPolynom;	//多项式系数组
		BigNrVec m_vSharingParts;	//share值向量
		vector<UInt> x_id_set;	//x_id向量组
		BigNumber secret;	//秘密；
	
};
