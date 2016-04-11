#include<iostream>
#include<fstream>
#include<string>
#include<vector>
#include<time.h>
#include <stdio.h>
#include <stdlib.h>
#include<bitset>

#include"ShamirSSScheme.h"
#include<string>
#include<vector>

using namespace std;

//将A文件进行Shamir's SSS分割成n份share;

void generatePrime(string filePath,UInt PRIME_LEN);
BigNumber getPrimeFromFile(string filePath);

int main(){

	vector<string> file_set;
	/*
	file_set.push_back("8450");
	file_set.push_back("8706");
	file_set.push_back("8962");
	file_set.push_back("9218");
	file_set.push_back("9474");
	file_set.push_back("9730");
	file_set.push_back("9986");*/
	file_set.push_back("258");

	vector<UInt> bit_set;
	/*
	bit_set.push_back(8450);
	bit_set.push_back(8706);
	bit_set.push_back(8962);
	bit_set.push_back(9218);
	bit_set.push_back(9474);
	bit_set.push_back(9730);
	bit_set.push_back(9986);*/
	bit_set.push_back(258);


	for (int i = 0; i < bit_set.size(); i++){
	string filePath;
	cout << "please input the name of file." << endl;
	filePath=file_set[i];
	generatePrime(filePath, bit_set[i]);
	BigNumber qq = getPrimeFromFile(filePath);
	cout << "qq:" << qq << endl;
}
//================================================================================================
	


	return 0;
}

void generatePrime(string filePath, UInt PRIME_LEN){
	UInt PRIME_LENGTH = PRIME_LEN;
	ofstream shuchu,explain;
	
	cout << "please input the length of prime." << endl;
	//cin >> PRIME_LENGTH;
	string explainPath = "time" + to_string(PRIME_LENGTH) + ".txt";
	explain.open(explainPath);
	explain << "PRIME_LENGTH:" << PRIME_LENGTH << endl;
	clock_t start, end;
	start = clock();
	BigNumber prime = NTL::GenPrime_ZZ(PRIME_LENGTH, 80);
	end = clock();
	double duration;
	duration = (double)(end - start) / CLOCKS_PER_SEC;
	cout << "生成素数运行时间：" << duration << endl;
	cout << "prime:" << prime << endl;
	explain << "The time of gererate of prime:" << duration << "s" << endl;
	vector<bitset<8> > bit_word_set((PRIME_LENGTH / 8)+1);
	bitset<8> lin;
	for (UInt i = 0; i < PRIME_LENGTH; i++){
		bit_word_set[i / 8][i % 8] = NTL::SwitchBit(prime, i);
	}
	string newFilePath = "F:\\360云盘\\code\\代码完成版\\素数\\SecretSSSchem\\SecretSSSchem\\素数\\"+filePath;
	ofstream outfile(newFilePath, ios::binary);
	if (!outfile){
		cerr << "open error!" << endl;
		abort();
	}
	for (int i = 0; i < bit_word_set.size(); i++){
		outfile.write((char*)& bit_word_set[i], 1);		//read size 以字节为单位读入；
	}
	outfile.close();
	shuchu.close();
	explain.close();
}
BigNumber getPrimeFromFile(string filePath){

	ifstream infile(filePath, ios::binary);
	if (!infile){
		cerr << "open error!" << endl;
		abort();
	}
	int length = 0;
	infile.seekg(0, ios::end);
	length = infile.tellg();

	infile.seekg(0, ios::beg);
	//申请length个word_set个空间；既申请被分成N个小块数；
	vector<bitset<8> > bit_word_set_1(length);
	for (int i = 0; i < length; i++){
		infile.read((char*)& bit_word_set_1[i], 1);		//read size 以字节为单位读入；
		//cout << "读入数据No"<<i<<":"<< bit_word_set[i] << endl;
	}

	infile.close();
	BigNumber rePrime;
	for (UInt i = 0; i < bit_word_set_1.size(); i++){
		bitset<8> linlin = bit_word_set_1[i];
		for (int j = 0; j < linlin.size(); j++){
			if (linlin[j])
			{
				NTL::SetBit(rePrime, i * 8 + j);
			}
		}
	}
	//BigNumber newprime = rePrime;
	cout << "reprime:" << rePrime << endl;
	cout << "reprime length:" << NTL::NumBits(rePrime) << endl;
	cout << "reprime sign:" << NTL::sign(rePrime) << endl;
	return rePrime;

}