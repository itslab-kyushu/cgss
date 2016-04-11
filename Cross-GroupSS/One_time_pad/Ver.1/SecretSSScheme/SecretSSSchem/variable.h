#pragma once
#pragma comment(lib,"D://code//tool//WinNTL-9_4_0//src//Debug//NTL.LIB")
#include<NTL//ZZ.h>
#include<vector>
#include<string>
using namespace std;

#define BIT256 256
#define BitLen_128 4096
#define BitLen_136 (BitLen_128 + 8)
#define PRIME_LENGTH (BitLen_128 + 2)
#define PRIME_FILE_NAME (BitLen_128 + 2)

typedef unsigned int UInt;
typedef NTL::ZZ BigNumber;
typedef std::vector<BigNumber> BigNrVec;

class Var{
	public: 
		void init();
		string SecretFilePathIn;
		string SecretFilePathOut;
		string ShareFileIn;
		string ShareFileOut;
		string PrimeFilePath;
		UInt TotalShares;
		UInt ThresholdShares;
		UInt TotalProviders;
		UInt ThresholdProviders;
		vector<UInt> X_ID;
		vector<UInt> Y_ID;
		string ResultFileOut;
		UInt Count;
		UInt DelCount;
		string KeyPath;

	private:
		void RecongizeCommand(string s);
};

