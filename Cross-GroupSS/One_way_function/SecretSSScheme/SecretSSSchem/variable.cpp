#include"variable.h"
#include<fstream>
#include<sstream>
//读取设置文件
void Var::init(){
	ifstream infile("InPutFile\\var.txt");
	if (!infile){
		cerr << "open error!" << endl;
		abort();
	}
	for (string s; getline(infile, s);){
		RecongizeCommand(s);
	}
}
//判断命令模式
void Var::RecongizeCommand(string s){
	string head(s, 0, s.find_first_of(":"));
	if (head == "SecretFilePathIn"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		SecretFilePathIn = ss;
	}
	if (head == "SecretFilePathOut"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		SecretFilePathOut = ss;
	}
	if (head == "ShareFilePathIn"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		ShareFileIn = ss;
	}
	if (head == "ShareFilePathOut"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		ShareFileOut = ss;
	}
	if (head == "PrimeFilePath"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		PrimeFilePath = ss;
	}
	if (head == "X_ID"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		for (UInt a; sin >> a;){
			X_ID.push_back(a);
		}
	}
	if (head == "Y_ID"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		for (UInt a; sin >> a;){
			Y_ID.push_back(a);
		}
	}
	if (head == "TotalShares"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		TotalShares = a;

	}
	if (head == "TotalProviders"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		TotalProviders = a;

	}
	if (head == "ThresholdShares"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		ThresholdShares = a;
	}
	if (head == "ThresholdProviders"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		ThresholdProviders = a;
	}
	if (head == "ResultFileOut"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		ResultFileOut = ss;
	}
	if (head == "Count"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		Count = a;
	}
	if (head == "DelCount"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		istringstream sin(ss);
		UInt a;
		sin >> a;
		DelCount = a;
	}
	if (head == "KeyPath"){
		string ss(s, s.find_first_of(":") + 1, s.length() - s.find_first_of(":") - 1);
		KeyPath = ss;
	}
}