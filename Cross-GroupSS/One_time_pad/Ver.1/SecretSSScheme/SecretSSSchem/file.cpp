#include"file.h"
#include<time.h>

void File::clear(){
	filePath = "";
	word_set.clear();
	modFileSize = NTL::to_ZZ(0);
	len = 0;
}

//写入文件电脑中的路径；
void File::write_filePath(string Path){
	filePath = Path;
}

//将文件流写入数组中；
void File::write_normal_file_into_word_set(string Path){

	filePath = Path;
	ifstream infile(filePath, ios::binary);
	if (!infile){
		cerr << "open error!" << endl;
		abort();
	}
	int length = 0;
	infile.seekg(0, ios::end);
	length = infile.tellg();
	infile.seekg(0, ios::beg);

	long long int modsize = length % (BitLen_128 / 8);
	if (modsize != 0){
		length = length / (BitLen_128 / 8) + 1;
	}
	else{
		length = length / (BitLen_128 / 8);
	}

	word_set.resize(length + 1);

	UInt blockSize = BitLen_128 / 8;
	for (int i = 0; i < length; i++){
		unsigned char zz[BitLen_128 / 8];
		infile.read((char*)& zz, blockSize);		//read size 以字节为单位读入；		
		NTL::ZZFromBytes(word_set[i], zz, blockSize);
	}

	infile.clear();
	infile.close();
	
	word_set[word_set.size() - 1] = modsize;

}
//将数组中BigNumber写入sharing文件流中
void File::write_word_set_into_share_file(string outPath){
	ofstream outfile(outPath, ios::binary);
	if (!outfile){
		cerr << "open error!" << endl;
		abort();
	}
	UInt WordSetSize = word_set.size();
	for (int i = 0; i < WordSetSize; i++){
		unsigned char zz[BitLen_136 / 8];
		NTL::BytesFromZZ(zz, word_set[i], BitLen_136 / 8);
		outfile.write((char*)& zz, (BitLen_136 / 8));
	}
	//关闭输出文件

	outfile.clear();
	outfile.close();

}


//将share文件流写入数组中；
void File::write_share_file_into_word_set(string Path){

	filePath = Path;
	ifstream infile(filePath, ios::binary);
	if (!infile){
		cerr << "open error!" << endl;
		abort();
	}
	int length = 0;
	infile.seekg(0, ios::end);
	length = infile.tellg();

	long long int modsize = length % (BitLen_136 / 8);

	if (modsize != 0){
		length = length / (BitLen_136 / 8) + 1;
	}
	else{
		length = length / (BitLen_136 / 8);
	}

	infile.seekg(0, ios::beg);
	//申请length个word_set个空间；既申请被分成N个小块数；
	word_set.resize(length);

	for (int i = 0; i < length; i++){
		unsigned char zz[BitLen_136 / 8];
		infile.read((char*)& zz, (BitLen_136 / 8));		//read size 以字节为单位读入；
		NTL::ZZFromBytes(word_set[i], zz, BitLen_136 / 8);
	}
	infile.clear();
	infile.close();
	
}
//将数组中BigNumber写入普通文件流中
void File::write_word_set_into_normal_file(string outPath){
	ofstream outfile(outPath, ios::binary);
	if (!outfile){
		cerr << "open error!" << endl;
		abort();
	}
	UInt WordSetSize = word_set.size();

	if (word_set[word_set.size() - 1] == 0){
		for (int i = 0; i < WordSetSize - 1; i++){
			unsigned char zz[BitLen_128 / 8];
			NTL::BytesFromZZ(zz, word_set[i], BitLen_128 / 8);
			outfile.write((char*)& zz, (BitLen_128 / 8));		//read size 以字节为单位读入；	
		}
	}
	else{
		for (int i = 0; i < WordSetSize - 2; i++){
			unsigned char zz[BitLen_128 / 8];
			NTL::BytesFromZZ(zz, word_set[i], BitLen_128 / 8);
			outfile.write((char*)& zz, (BitLen_128 / 8));		//read size 以字节为单位读入；
		}
		//unsigned char cc[];
		unsigned char bb[BitLen_128 / 8];
		NTL::BytesFromZZ(bb, word_set[WordSetSize - 2], BitLen_128 / 8);
		BigNumber modsize = word_set[WordSetSize - 1];
		cout << "modsize:" << modsize << endl;
		outfile.write((char*)& bb, NTL::to_uint(modsize));

	}
	//关闭输出文件

	outfile.clear();
	outfile.close();

}

//获取word_set();
const vector<BigNumber> & File::get_word_set(){
	return word_set;
}
//重置word_set大小
void File::resize_word_set(UInt n){
	word_set.resize(n);
}
//将一个向量数据写到word_set中
void File::write_vec_into_word_set(vector<BigNumber> &vv){
	word_set = vv;
}

//将单独一个数据写到word_set的i号位置上
void File::write_vec_i_into_word_set(BigNumber data, unsigned int i){
	word_set[i] = data;
}
//将单独一个数据写到word_set的末尾
void File::write_current_push_back_word_set(BigNumber cur){
	word_set.push_back(cur);
}
//从word_set的i号位置读出数据
const BigNumber File::get_posI_from_word_set(unsigned int i){
	return word_set[i];
}

//将写入modFileSize中
void File::write_modFileSize(BigNumber bn){
	modFileSize = bn;
}
//读出modFileSize
const BigNumber File::read_modFileSize(){
	return modFileSize;
}
const string File::get_filePath(){
	return filePath;
}