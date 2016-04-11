#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<vector>
#include "variable.h"

using namespace std;

class ShamirSSScheme{
	
	public:

		void init(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET);		//��ʼ������;	
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, BigNrVec& sharePart);	//���ò���1
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime);	//���ò���2
		void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, bool whetherGenPoly);//���ò���3
		void write_m_nK(UInt nK);						//д��ֵ
		void write_m_nN(UInt nN);						//дshare����ֵ
		void write_m_bnPrimeNr(BigNumber m_bnPrime);	//д������ֵ
		void write_x_id_set(vector<UInt>& x_id_set);		//дx_id����ֵ 
		void write_x_id_set_i(UInt i, UInt data);		//дx_id��i��λ��ֵ
		void write_m_vSharingPart(BigNrVec& sharePart);	//д��sharePart
		void write_secret(BigNumber s);							//д��secret;


		const BigNumber get_m_bnPrimeNr();		//ȡ������ֵ
		const vector<UInt>& get_x_id_set();		//ȡx_id����
		const UInt get_x_id_set_i(UInt i);		//ȡx_id��i��λ��ֵ
		const BigNrVec& get_m_vSharing();		//ȡsharePart;
		const BigNumber get_secret();					//ȡsecret;

		void gen_Prime();				//���ɴ�����
		void gen_Polynom();				//���ɶ���ʽ
		void gen_SecretSharePart();		//����share
		void secret_Reconstruction();   //���¹�������

	private:
		UInt m_nN; //share����
		UInt m_nK; //��ֵ
		BigNumber m_bnPrimeNr; //������
		BigNrVec m_vPolynom;	//����ʽϵ����
		BigNrVec m_vSharingParts;	//shareֵ����
		vector<UInt> x_id_set;	//x_id������
		BigNumber secret;	//���ܣ�
	
};
