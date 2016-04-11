#pragma once
#include<vector>
#include<string>
#include<bitset>
#include<fstream>
#include "variable.h"

using namespace std;

//�����Լ�����T�Ľ������ݶ�ȡ����������T���Ըı��д�����ֽڳ���

class File
{
public:

	//д���ļ������е�·����
	void write_filePath(string Path);
	//����ͨ�ļ���д�������У�
	void write_normal_file_into_word_set(string Path);
	//��share�ļ�д�������У�
	void write_share_file_into_word_set(string Path);
	//��������BigNumberд����ͨ�ļ�����
	void write_word_set_into_normal_file(string outPath);
	//��������BigNumberд��sharing�ļ�����
	void write_word_set_into_share_file(string outPath);
	//��ȡword_set();
	const vector<BigNumber> & get_word_set();
	//����word_set��С
	void resize_word_set(UInt n);
	//��һ����������д��word_set��
	void write_vec_into_word_set(vector<BigNumber> &vv);
	//������һ������д��word_set��i��λ����
	void write_vec_i_into_word_set(BigNumber data, unsigned int i);
	//������һ������д��word_set��ĩβ
	void write_current_push_back_word_set(BigNumber cur);
	//��word_set��i��λ�ö�������
	const BigNumber get_posI_from_word_set(unsigned int i);
	//��д��modFileSize��
	void write_modFileSize(BigNumber bn);
	//����modFileSize
	const BigNumber read_modFileSize();
	void clear();

private:

	//filePath��ʾ�ļ�·����
	string filePath;	
	//word_set��ʾ�����洢���Լ��ļ��е���������
	vector<BigNumber> word_set;
	//�������һ������˼����ֽڣ�
	BigNumber modFileSize;
	//bitset�ĳ���
	UInt len;
};

