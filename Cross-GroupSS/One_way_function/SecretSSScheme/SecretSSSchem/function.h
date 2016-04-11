#include<iostream>
#include"file.h"
#include<vector>
#include<algorithm>
#include<time.h>
#include"ShamirSSScheme.h"
#include"variable.h"
#include"key.h"

#define KeySize 256

using namespace std;

void inPutNK(unsigned int& n, unsigned int& k);
void inPutMT(unsigned int& m, unsigned int& t);
void inPutX_ID(vector<UInt>& x_id);
void inPutY_ID(vector<UInt>& y_id);
bool judgeLegal(int n, int k, int m, int t);
void outPutTime(clock_t start, clock_t end, string typeword, vector<double>& time);

vector<File> fileShamirSSS_Sharing(File& needSharingFile, UInt n, UInt k, vector<UInt> x_id, BigNumber bigPrime);
File fileShamirSSS_Reconstruction(vector<File>&fileSet, UInt n, UInt k, vector<UInt> x_id, BigNumber bigPrime);

void readFile(File& file, string filePath);
void sMinusQ(File& file, BigNumber q);
void pPlusQ(File& file, BigNumber q_v);
void outPutFile(File& file, string outPath);
void secretSharing(vector<File>& FileSet, File& bigFile, UInt n, UInt k, vector<UInt>& x_id, BigNumber prime);
void outPutshareFiles(vector<File>& fileSet);
void outPutAuthorityFiles(vector<BigNumber>& fileSet);
void readShareFiles(vector<File>& shareInFile);
void secretReconstruction_V(ShamirSSScheme& ReAuthorityTool, UInt m, UInt t, BigNumber prime_Au, vector<BigNumber>& AuthorityShare_Set, vector<UInt>& y_id);
void genAuShares(vector<BigNumber>& AuthorityShare_Set, BigNumber v, UInt t, UInt m, vector<UInt>& y_id, BigNumber& prime_Au, vector<double>&time_v_share);
BigNumber getPrimeFromFile(string filePath);
void outputResult(ofstream & shuchu, string message, vector<double> & set, Var& varia);
void outputResult(ofstream & shuchu, ofstream & sabun, string message, vector<double> & set, Var& varia);
void readAuthrotiy(vector<BigNumber>& FileAuthShare_Set, Var& varia);
void outPutKey(BigNumber key, Var& varia);
BigNumber readKey(Var& varia);
void outSecKey(SecByteBlock h_sec,Var& varia);
SecByteBlock readSecKey(Var& varia);