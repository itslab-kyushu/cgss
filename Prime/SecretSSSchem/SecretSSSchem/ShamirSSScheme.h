#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<vector>
using namespace std;
#define BitLen_128 1792
typedef unsigned int UInt;
typedef NTL::ZZ BigNumber;
typedef std::vector<BigNumber> BigNrVec;
class ShamirSSScheme{
	
	public:

		void init(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET);		//初始化程序;	
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, BigNrVec& sharePart);	//配置参数1
		void set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime);	//配置参数2
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

//初始化程序;写入阈值，share总数，秘密，x_id;
void ShamirSSScheme::init(unsigned int NK, unsigned int NN, BigNumber S, vector<unsigned int>& X_ID_SET){
	m_vSharingParts.resize(NN);
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	gen_Prime();         //生成大素数之后才可以生成多项式
	gen_Polynom();		 //生成多项式
}

//配置参数1
void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, BigNrVec& sharePart){
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	write_m_bnPrimeNr(prime);
	write_m_vSharingPart(sharePart);
}

//配置参数2
void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime){
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	write_m_bnPrimeNr(prime);
}

//写阈值
void ShamirSSScheme::write_m_nK(unsigned int nK){
	m_nK = nK;
}
//写share总数值
void ShamirSSScheme::write_m_nN(unsigned int nN){
	m_nN = nN;
}
//写大素数值
void ShamirSSScheme::write_m_bnPrimeNr(NTL::ZZ m_bnPrime){
	m_bnPrimeNr = m_bnPrime;
}
//写x_id向量值
void ShamirSSScheme::write_x_id_set(vector<UInt>& x_id){
	x_id_set = x_id;
}
//写x_id的i号位置值
void ShamirSSScheme::write_x_id_set_i(UInt i, UInt data){
	x_id_set[i] = data;
}
//写入sharePart
void ShamirSSScheme::write_m_vSharingPart(BigNrVec& sharePart){
	m_vSharingParts = sharePart;
}
//写入secret;
void ShamirSSScheme::write_secret(BigNumber s){
	secret = s;
}

//取大素数值
const NTL::ZZ ShamirSSScheme::get_m_bnPrimeNr(){
	return m_bnPrimeNr;
}
//取x_id向量
const vector<unsigned int>& ShamirSSScheme::get_x_id_set(){
	return x_id_set;
}
//取x_id的i号位置值
const unsigned int ShamirSSScheme::get_x_id_set_i(UInt i){
	return x_id_set[i];
}

const vector<NTL::ZZ>& ShamirSSScheme::get_m_vSharing(){
	return m_vSharingParts;
}
//取secret;
const BigNumber ShamirSSScheme::get_secret(){
	return secret;
}

//生成大素数
void ShamirSSScheme::gen_Prime(){
	m_bnPrimeNr = NTL::GenPrime_ZZ(BitLen_128, 80);
}
//生成多项式
void ShamirSSScheme::gen_Polynom(){
	m_vPolynom.resize(m_nK);//重置数组大小；
	for (UInt i = 1; i < m_nK; i++){
		NTL::RandomBnd(m_vPolynom[i], m_bnPrimeNr);
	}
}
//生成share
void ShamirSSScheme::gen_SecretSharePart(){
	BigNumber aux;
	BigNumber Lin_m_vSharingPart;
	m_vSharingParts.resize(m_nN);
	for (UInt i = 0; i < m_nN; i++){
		Lin_m_vSharingPart = NTL::to_ZZ(secret); //首数；
		//cout << "Lin_m_vSharingPart:" << Lin_m_vSharingPart << endl;
		for (UInt j = 1; j < m_nK; j++){
			NTL::PowerMod(aux, NTL::to_ZZ(x_id_set[i]), j, m_bnPrimeNr);
			Lin_m_vSharingPart = (Lin_m_vSharingPart + m_vPolynom[j] * aux) % m_bnPrimeNr;
		}
		m_vSharingParts[i] = Lin_m_vSharingPart; //写入保存share值数组中；
		//cout << "pre_secret:" << Lin_m_vSharingPart << endl;
	}

	
}		

//重新构造秘密
void ShamirSSScheme::secret_Reconstruction(){
	
	BigNumber or_secret = NTL::to_ZZ(secret), aux, aux1;
	//cout << "or=secret"<<or_secret << endl;
	for (UInt i = 0; i < m_vSharingParts.size(); i++){
		aux1 = 1;
		for (UInt j = 0; j < m_vSharingParts.size(); j++){
			if (x_id_set[j] != x_id_set[i]){
				aux = x_id_set[j] - x_id_set[i];
				while (aux<=0)
				{
					aux += m_bnPrimeNr;
				}
				NTL::MulMod(aux1, aux1, (x_id_set[j]*NTL::InvMod(aux, m_bnPrimeNr)) % m_bnPrimeNr, m_bnPrimeNr);
			}
		}
			or_secret = (or_secret + m_vSharingParts[i]*aux1) % m_bnPrimeNr;
	}


	BigNumber aa;
	NTL::SetBit(aa, BitLen_128+1);

	if (or_secret >= aa){
		or_secret = or_secret - m_bnPrimeNr;
		cout << "in:" << or_secret << endl;
	}
	//cout << "or-hou:" << NTL::NumBits(or_secret) << endl;
	secret = or_secret;
	//cout << "finall_or_secret2：" << secret << endl;
}
