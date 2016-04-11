#include"key.h"
#include<bitset>
#include<string>
#define KeySize 256
//生成一个256位的权限秘密
string Key::genKey(){
	AutoSeededRandomPool prng;
	SecByteBlock key(KeySize / 8);
	prng.GenerateBlock(key, key.size());
	string encoded;
	encoded.clear();														// HexEncoder
	StringSource ss1(key, key.size(), true, new HexEncoder(new StringSink(encoded))); // StringSource
	return encoded;
}
SecByteBlock Key::genKey_Sec(){
	AutoSeededRandomPool prng;
	SecByteBlock key(KeySize / 8);
	prng.GenerateBlock(key, key.size());
	return key;
}
string Key::covSecKeyToString(SecByteBlock key){
	SecByteBlock keynew = key;
	string encoded;
	encoded.clear();														// HexEncoder
	StringSource ss1(keynew, keynew.size(), true, new HexEncoder(new StringSink(encoded))); // StringSource
	return encoded;
}

string Key::covHextoBinary(string Hex){
	string str_binary;
	for (int i = 0; i<Hex.size(); i++){
		switch (Hex[i])
		{
		case '0': str_binary = str_binary + "0000"; break;
		case '1': str_binary = str_binary + "0001"; break;
		case '2': str_binary = str_binary + "0010"; break;
		case '3': str_binary = str_binary + "0011"; break;
		case '4': str_binary = str_binary + "0100"; break;
		case '5': str_binary = str_binary + "0101"; break;
		case '6': str_binary = str_binary + "0110"; break;
		case '7': str_binary = str_binary + "0111"; break;
		case '8': str_binary = str_binary + "1000"; break;
		case '9': str_binary = str_binary + "1001"; break;
		case 'A': str_binary = str_binary + "1010"; break;
		case 'B': str_binary = str_binary + "1011"; break;
		case 'C': str_binary = str_binary + "1100"; break;
		case 'D': str_binary = str_binary + "1101"; break;
		case 'E': str_binary = str_binary + "1110"; break;
		case 'F': str_binary = str_binary + "1111"; break;
		default:
			break;
		}
	}
	return str_binary;
}
NTL::ZZ Key::covHextoBigNumber(string Hex){
	string str_binary = "";
	string revise_Hex = Hex;
	for (int i = 0; i<revise_Hex.size(); i++){
		switch (revise_Hex[i])
		{
		case '0': str_binary = str_binary + "0000"; break;
		case '1': str_binary = str_binary + "0001"; break;
		case '2': str_binary = str_binary + "0010"; break;
		case '3': str_binary = str_binary + "0011"; break;
		case '4': str_binary = str_binary + "0100"; break;
		case '5': str_binary = str_binary + "0101"; break;
		case '6': str_binary = str_binary + "0110"; break;
		case '7': str_binary = str_binary + "0111"; break;
		case '8': str_binary = str_binary + "1000"; break;
		case '9': str_binary = str_binary + "1001"; break;
		case 'A': str_binary = str_binary + "1010"; break;
		case 'B': str_binary = str_binary + "1011"; break;
		case 'C': str_binary = str_binary + "1100"; break;
		case 'D': str_binary = str_binary + "1101"; break;
		case 'E': str_binary = str_binary + "1110"; break;
		case 'F': str_binary = str_binary + "1111"; break;
		default:
			break;
		}
	}
	NTL::ZZ bigNumber;
	bigNumber = NTL::to_ZZ(0);
	for (int i = 0; i < str_binary.size(); i++){
		if (str_binary[i] == '1')
			NTL::SetBit(bigNumber, str_binary.size() - 1 - i);
	}
	return bigNumber;
}

//caculate q=fh(v);
NTL::ZZ Key::genOneTimePad(NTL::ZZ v, NTL::ZZ r){
	NTL::ZZ bigNumber;
	NTL::bit_xor(bigNumber, v, r);		//进行异或操作
	return bigNumber;
}
//------------------------------------------------------------------------
//データのタイプはBigNumberからStringに変更する; StringはHexで表示する；
//------------------------------------------------------------------------
string Key::covBigNumberToString(NTL::ZZ bigN){
	NTL::ZZ BigN = bigN;
	string returnStr;
	int bitBigN = NTL::NumBits(BigN);
	if (bitBigN % 4 != 0){
		bitBigN = bitBigN / 4 + 1;
	}
	else{
		bitBigN = bitBigN / 4;
	}

	vector<bitset<4> > sinkPool(bitBigN);
	for (int i = 0; i < NTL::NumBits(BigN); i++){
		sinkPool[i / 4][i % 4] = NTL::SwitchBit(BigN, i);
	}
	for (int i = 0; i<sinkPool.size(); i++){
		unsigned int switch_on = sinkPool[i].to_ulong();
		switch (switch_on)
		{
		case 0:  returnStr = "0" + returnStr; break;
		case 1:  returnStr = "1" + returnStr; break;
		case 2:  returnStr = "2" + returnStr; break;
		case 3:  returnStr = "3" + returnStr; break;
		case 4:  returnStr = "4" + returnStr; break;
		case 5:  returnStr = "5" + returnStr; break;
		case 6:  returnStr = "6" + returnStr; break;
		case 7:  returnStr = "7" + returnStr; break;
		case 8:  returnStr = "8" + returnStr; break;
		case 9:  returnStr = "9" + returnStr; break;
		case 10: returnStr = "A" + returnStr; break;
		case 11: returnStr = "B" + returnStr; break;
		case 12: returnStr = "C" + returnStr; break;
		case 13: returnStr = "D" + returnStr; break;
		case 14: returnStr = "E" + returnStr; break;
		case 15: returnStr = "F" + returnStr; break;
		default:
			break;
		}
	}
//-----------------------------------------------------------------
// 补足由于bigNumber转化成string可能造成的高位少0现象
//-----------------------------------------------------------------
	if (returnStr.size() < 64){
		unsigned int count = 64 - returnStr.size();
		for (int j = 0; j <count; j++){
			returnStr = "0" + returnStr;

		}
	}
	return returnStr;
}
//生成一个长度为l bit的随机数
NTL::ZZ Key::genR(unsigned int l){
	NTL::ZZ r;
	NTL::ZZ big;
	NTL::SetBit(big, l);	//generate the length l+1 number;
	NTL::RandomBnd(r, big);	//generate random number between 0~2^l
	return r;
}