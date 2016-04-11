#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<bitset>
#include "variable.h"
//����ת������С��


using namespace std;

typedef unsigned int UInt;
typedef NTL::ZZ BigNumber;

class converter{

	public: 
		const bitset<BitLen_128> BigNum_To_BitSet_128(BigNumber BigNum);			//��BigNumber��ת����BitSet���ͣ�
		const bitset<BitLen_136> BigNum_To_BitSet_136(BigNumber BigNum);			//��BigNumber��ת����BitSet���ͣ�
		const BigNumber BitSet_To_BigNum_128(bitset<BitLen_128> BitSet);			//��BitSet����ת����BigNumber�ͣ�
		const BigNumber BitSet_To_BigNum_136(bitset<BitLen_136> BitSet);			//��BitSet����ת����BigNumber�ͣ�
	private:
		bitset<BitLen_128> bitBigNum128;	//��Ŵ�����bit������bit�����Ĵ�СΪT���ֽڣ�128λ��
		bitset<BitLen_136> bitBigNum136;	//��Ŵ�����bit������bit�����Ĵ�СΪT���ֽڣ�136λ��
		BigNumber bigNum;				//���������ȿ��Ը��ģ�

};
//��BigNumber��ת����BitSet���ͣ�
const bitset<BitLen_128> converter::BigNum_To_BitSet_128(BigNumber BigNum){
	
	bitset<BitLen_128> lin;
	unsigned int n = NTL::NumBits(BigNum);
	for (int i = 0; i < n; i++){
		lin[i] = NTL::SwitchBit(BigNum, i);
	}
	bitBigNum128 = lin;
	return bitBigNum128;
}
//��BigNumber��ת����BitSet���ͣ�
const bitset<BitLen_136> converter::BigNum_To_BitSet_136(BigNumber BigNum){

	bitset<BitLen_136> lin;
	unsigned int n = NTL::NumBits(BigNum);
	for (int i = 0; i < n; i++){
		lin[i] = NTL::SwitchBit(BigNum, i);
	}
	bitBigNum136 = lin;
	return bitBigNum136;
}
//��BitSet����ת����BigNumber�ͣ�
const BigNumber converter::BitSet_To_BigNum_128(bitset<BitLen_128> BitSet){
	UInt n = BitSet.size();
	BigNumber lin=NTL::to_ZZ(0);
	for (int i = 0; i < n; i++){
		if (BitSet[i])
			NTL::SetBit(lin, i);
	}
	bigNum = lin;
	return bigNum;
}
//��BitSet����ת����BigNumber�ͣ�
const BigNumber converter::BitSet_To_BigNum_136(bitset<BitLen_136> BitSet){
	UInt n = BitSet.size();
	BigNumber lin = NTL::to_ZZ(0);
	for (int i = 0; i < n; i++){
		if (BitSet[i])
			NTL::SetBit(lin, i);
	}
	bigNum = lin;
	return bigNum;
}