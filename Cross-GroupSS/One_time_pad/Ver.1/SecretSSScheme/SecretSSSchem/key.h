#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include <cstdio>
#include <iostream>
#include <string>
#include "..\Crypto\osrng.h"
#include"..\Crypto\osrng.h"
#include"..\Crypto\sha.h"
#include"..\Crypto\filters.h"
#include"..\Crypto\hmac.h"
#include"..\Crypto\hex.h"
#include"NTL\ZZ.h"

using namespace CryptoPP;
using namespace std;

class Key
{
public:
	string genKey();												//generate key; The key type is string;The key can be as secret key h or authority key v;
	SecByteBlock genKey_Sec();										//generate key; the key type is SecByteBlock;
	string covSecKeyToString(SecByteBlock key);						//transfer SecByteBlock to string;
	SecByteBlock covStringKeyToSec(string key);						//transfer string to SecByteBlock;		(Î´Íê³É)
	string covHextoBinary(string Hex);								//transfer hex to binary;
	NTL::ZZ covHextoBigNumber(string Hex);							//transfer hex to bigNumber;
	NTL::ZZ genOneTimePad(NTL::ZZ v, NTL::ZZ r);	//caculate q=fh(v); Both secretKey_h and authority are under Hex model;
	string covBigNumberToString(NTL::ZZ bigN);						// Transfer bigNumber to Hex string
	NTL::ZZ genR(unsigned int l);											//Generate a random one time pad, and the length is l

private:

};
