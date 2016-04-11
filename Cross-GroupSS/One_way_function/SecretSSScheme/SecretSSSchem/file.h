#pragma once
#include<vector>
#include<string>
#include<bitset>
#include<fstream>
#include "variable.h"

using namespace std;

//根据自己定义T的进入数据读取（数据类型T可以改变读写数据字节长）

class File
{
public:

	//写入文件电脑中的路径；
	void write_filePath(string Path);
	//将普通文件流写入数组中；
	void write_normal_file_into_word_set(string Path);
	//将share文件写入数组中；
	void write_share_file_into_word_set(string Path);
	//将数组中BigNumber写入普通文件流中
	void write_word_set_into_normal_file(string outPath);
	//将数组中BigNumber写入sharing文件流中
	void write_word_set_into_share_file(string outPath);
	//获取word_set();
	const vector<BigNumber> & get_word_set();
	//重置word_set大小
	void resize_word_set(UInt n);
	//将一个向量数据写到word_set中
	void write_vec_into_word_set(vector<BigNumber> &vv);
	//将单独一个数据写到word_set的i号位置上
	void write_vec_i_into_word_set(BigNumber data, unsigned int i);
	//将单独一个数据写到word_set的末尾
	void write_current_push_back_word_set(BigNumber cur);
	//从word_set的i号位置读出数据
	const BigNumber get_posI_from_word_set(unsigned int i);
	//将写入modFileSize中
	void write_modFileSize(BigNumber bn);
	//读出modFileSize
	const BigNumber read_modFileSize();
	void clear();

private:

	//filePath表示文件路径；
	string filePath;	
	//word_set表示用来存储来自己文件中的数据流；
	vector<BigNumber> word_set;
	//用来最后一组读进了几个字节；
	BigNumber modFileSize;
	//bitset的长度
	UInt len;
};

