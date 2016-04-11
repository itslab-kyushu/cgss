#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<bitset>
#include "variable.h"
//定义转换器大小；


using namespace std;

typedef unsigned int UInt;
typedef NTL::ZZ BigNumber;

class converter{

	public: 
		const bitset<BitLen_128> BigNum_To_BitSet_128(BigNumber BigNum);			//从BigNumber型转化成BitSet类型；
		const bitset<BitLen_136> BigNum_To_BitSet_136(BigNumber BigNum);			//从BigNumber型转化成BitSet类型；
		const BigNumber BitSet_To_BigNum_128(bitset<BitLen_128> BitSet);			//从BitSet类型转化成BigNumber型；
		const BigNumber BitSet_To_BigNum_136(bitset<BitLen_136> BitSet);			//从BitSet类型转化成BigNumber型；
	private:
		bitset<BitLen_128> bitBigNum128;	//存放大数的bit向量，bit向量的大小为T个字节，128位；
		bitset<BitLen_136> bitBigNum136;	//存放大数的bit向量，bit向量的大小为T个字节，136位；
		BigNumber bigNum;				//大数，长度可以更改；

};
//从BigNumber型转化成BitSet类型；
const bitset<BitLen_128> converter::BigNum_To_BitSet_128(BigNumber BigNum){
	
	bitset<BitLen_128> lin;
	unsigned int n = NTL::NumBits(BigNum);
	for (int i = 0; i < n; i++){
		lin[i] = NTL::SwitchBit(BigNum, i);
	}
	bitBigNum128 = lin;
	return bitBigNum128;
}
//从BigNumber型转化成BitSet类型；
const bitset<BitLen_136> converter::BigNum_To_BitSet_136(BigNumber BigNum){

	bitset<BitLen_136> lin;
	unsigned int n = NTL::NumBits(BigNum);
	for (int i = 0; i < n; i++){
		lin[i] = NTL::SwitchBit(BigNum, i);
	}
	bitBigNum136 = lin;
	return bitBigNum136;
}
//从BitSet类型转化成BigNumber型；
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
//从BitSet类型转化成BigNumber型；
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