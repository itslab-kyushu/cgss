#include"ShamirSSScheme.h"

//��ʼ������;д����ֵ��share���������ܣ�x_id;
void ShamirSSScheme::init(unsigned int NK, unsigned int NN, BigNumber S, vector<unsigned int>& X_ID_SET){
	m_vSharingParts.resize(NN);
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	gen_Prime();         //���ɴ�����֮��ſ������ɶ���ʽ
	gen_Polynom();		 //���ɶ���ʽ
}

//���ò���1
void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, BigNrVec& sharePart){
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	write_m_bnPrimeNr(prime);
	write_m_vSharingPart(sharePart);
}

//���ò���2
void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime){
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	write_m_bnPrimeNr(prime);
}
//���ò���3
void ShamirSSScheme::set_Parameter(UInt NK, UInt NN, BigNumber S, vector<UInt>& X_ID_SET, BigNumber prime, bool whetherGenPoly){
	write_m_nK(NK);
	write_m_nN(NN);
	write_x_id_set(X_ID_SET);
	write_secret(S);
	write_m_bnPrimeNr(prime);
	if (whetherGenPoly == true){
		gen_Polynom();
	}
}

//д��ֵ
void ShamirSSScheme::write_m_nK(unsigned int nK){
	m_nK = nK;
}
//дshare����ֵ
void ShamirSSScheme::write_m_nN(unsigned int nN){
	m_nN = nN;
}
//д������ֵ
void ShamirSSScheme::write_m_bnPrimeNr(NTL::ZZ m_bnPrime){
	m_bnPrimeNr = m_bnPrime;
}
//дx_id����ֵ
void ShamirSSScheme::write_x_id_set(vector<UInt>& x_id){
	x_id_set = x_id;
}
//дx_id��i��λ��ֵ
void ShamirSSScheme::write_x_id_set_i(UInt i, UInt data){
	x_id_set[i] = data;
}
//д��sharePart
void ShamirSSScheme::write_m_vSharingPart(BigNrVec& sharePart){
	m_vSharingParts = sharePart;
}
//д��secret;
void ShamirSSScheme::write_secret(BigNumber s){
	secret = s;
}

//ȡ������ֵ
const NTL::ZZ ShamirSSScheme::get_m_bnPrimeNr(){
	return m_bnPrimeNr;
}
//ȡx_id����
const vector<unsigned int>& ShamirSSScheme::get_x_id_set(){
	return x_id_set;
}
//ȡx_id��i��λ��ֵ
const unsigned int ShamirSSScheme::get_x_id_set_i(UInt i){
	return x_id_set[i];
}

const vector<NTL::ZZ>& ShamirSSScheme::get_m_vSharing(){
	return m_vSharingParts;
}
//ȡsecret;
const BigNumber ShamirSSScheme::get_secret(){
	return secret;
}

//���ɴ�����
void ShamirSSScheme::gen_Prime(){
	m_bnPrimeNr = NTL::GenPrime_ZZ(PRIME_LENGTH, 80);
}
//���ɶ���ʽ
void ShamirSSScheme::gen_Polynom(){
	m_vPolynom.resize(m_nK);//���������С��
	for (UInt i = 1; i < m_nK; i++){
		NTL::RandomBnd(m_vPolynom[i], m_bnPrimeNr);
	}
}
//����share
void ShamirSSScheme::gen_SecretSharePart(){
	BigNumber aux;
	BigNumber Lin_m_vSharingPart;
	m_vSharingParts.resize(m_nN);
	for (UInt i = 0; i < m_nN; i++){
		Lin_m_vSharingPart = NTL::to_ZZ(secret); //������
		//cout << "Lin_m_vSharingPart:" << Lin_m_vSharingPart << endl;
		for (UInt j = 1; j < m_nK; j++){
			NTL::PowerMod(aux, NTL::to_ZZ(x_id_set[i]), j, m_bnPrimeNr);
			Lin_m_vSharingPart = (Lin_m_vSharingPart + m_vPolynom[j] * aux) % m_bnPrimeNr;
		}
		m_vSharingParts[i] = Lin_m_vSharingPart; //д�뱣��shareֵ�����У�
		//cout << "pre_secret:" << Lin_m_vSharingPart << endl;
	}


}

//���¹�������
void ShamirSSScheme::secret_Reconstruction(){

	BigNumber or_secret = NTL::to_ZZ(secret), aux, aux1;
	//cout << "or=secret"<<or_secret << endl;
	for (UInt i = 0; i < m_vSharingParts.size(); i++){
		aux1 = 1;
		for (UInt j = 0; j < m_vSharingParts.size(); j++){
			if (x_id_set[j] != x_id_set[i]){
				aux = x_id_set[j] - x_id_set[i];
				while (aux <= 0)
				{
					aux += m_bnPrimeNr;
				}
				NTL::MulMod(aux1, aux1, (x_id_set[j] * NTL::InvMod(aux, m_bnPrimeNr)) % m_bnPrimeNr, m_bnPrimeNr);
			}
		}
		or_secret = (or_secret + m_vSharingParts[i] * aux1) % m_bnPrimeNr;
	}


	BigNumber aa;
	NTL::SetBit(aa, BitLen_128 + 1);

	if (or_secret >= aa){
		or_secret = or_secret - m_bnPrimeNr;
		cout << "in:" << or_secret << endl;
	}
	//cout << "or-hou:" << NTL::NumBits(or_secret) << endl;
	secret = or_secret;
	//cout << "finall_or_secret2��" << secret << endl;
}
