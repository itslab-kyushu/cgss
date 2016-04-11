#include"function.h"


void inPutNK(unsigned int& n, unsigned int& k){
	cout << "������n,k" << endl;
	cin >> n >> k;
}
void inPutMT(unsigned int& m, unsigned int& t){
	cout << "������m,t" << endl;
	cin >> m >> t;
}
void inPutX_ID(vector<UInt>& x_id){
	for (int i = 0; i < x_id.size(); i++){
		cout << "������No" << i << ":" << endl;
		cin >> x_id[i];
	}
}
void inPutY_ID(vector<UInt>& y_id){
	for (int i = 0; i < y_id.size(); i++){
		cout << "������ץ�Х�����IDNo" << i << ":" << endl;
		cin >> y_id[i];
	}
}
bool judgeLegal(int n, int k, int m, int t){
	if (n < m){
		cout << "Share������С����Ӫ�����ݣ��������" << endl;
		return false;
	}
	if (k < t){
		cout << "Share����ֵС����ӪӪ��ֵ���������" << endl;
		return false;
	}
	return true;
}
void outPutTime(clock_t start, clock_t end, string typeword, vector<double>& time){
	double duration;
	duration = (double)(end - start) / CLOCKS_PER_SEC;
	time.push_back(duration);

}
void readFile(File& file, string filePath){
	file.write_normal_file_into_word_set(filePath);
}

void outPutFile(File& file, string outPath){
	//string outPutFilePath; outPutFilePath = "testCopy.docx";
	file.write_word_set_into_normal_file(outPath);
}
//��A�ļ�����Shamir's SSS�ָ��n��share;
vector<File> fileShamirSSS_Sharing(File& needSharingFile, UInt n, UInt k, vector<UInt> x_id, BigNumber bigPrime){
	//share�ļ�
	vector<File> fileSet(n);
	//��ʱword_set�����洢A�ļ��е�word_set����ʱ���ݣ�
	vector<BigNumber> lin_word_set;
	lin_word_set = needSharingFile.get_word_set();
	//����һ����������Shamir's SSS�Ĺ��߶���e;
	ShamirSSScheme e;
	for (int i = 0; i < lin_word_set.size(); i++){
		e.set_Parameter(k, n, lin_word_set[i], x_id, bigPrime);
		e.gen_Polynom();					//���ɶ���ʽϵ��
		e.gen_SecretSharePart();			//����shares
		//����һ�������洢m_vSharing����ʱ����
		vector<BigNumber> lin_shares = e.get_m_vSharing();
		for (int j = 0; j < n; j++){
			fileSet[j].write_current_push_back_word_set(lin_shares[j]);//�ַ�lin_share�е�sharing��sharing�ļ��С�
		}
	}
	return fileSet;//����sharing�ļ�������
}

File fileShamirSSS_Reconstruction(vector<File>&fileSet, UInt n, UInt k, vector<UInt> x_id, BigNumber bigPrime){
	File bigFile;		//�ع�����ļ�
	vector<File> lin_fileSet = fileSet;	//��ȡsharing�ļ���
	vector<BigNumber> lin_word_set(fileSet.size());		//�ع�����ļ�����

	UInt loop = lin_fileSet[0].get_word_set().size();
	for (int i = 0; i < loop; i++){
		ShamirSSScheme e;
		e.write_m_nK(k);
		e.write_m_nN(n);
		e.write_x_id_set(x_id);
		e.write_m_bnPrimeNr(bigPrime);
		for (int j = 0; j < fileSet.size(); j++){
			lin_word_set[j] = lin_fileSet[j].get_posI_from_word_set(i);
		}
		e.write_m_vSharingPart(lin_word_set);
		e.secret_Reconstruction();
		bigFile.write_current_push_back_word_set(e.get_secret());
	}
	return bigFile;
}


void sMinusQ(File& file, BigNumber q){
	unsigned int n = file.get_word_set().size();
	for (int i = 0; i < n; i++){
		file.write_vec_i_into_word_set(file.get_posI_from_word_set(i) - q, i);
	}

}
void pPlusQ(File& file, BigNumber q_v){
	unsigned int n = file.get_word_set().size();
	for (int i = 0; i < n; i++){
		file.write_vec_i_into_word_set(file.get_posI_from_word_set(i) + q_v, i);
	}
}

void secretSharing(vector<File>& FileSet, File& bigFile, UInt n, UInt k, vector<UInt>& x_id, BigNumber prime){
	FileSet = fileShamirSSS_Sharing(bigFile, n, k, x_id, prime);			//����sharing �ļ���
}
void outPutshareFiles(vector<File>& fileSet){
	//���n��BigNumber�ļ�	
	for (int i = 0; i < fileSet.size(); i++){
		string SharesPath = "Share\\Sharing";
		SharesPath = SharesPath + to_string(i);
		fileSet[i].write_word_set_into_share_file(SharesPath);
	}
}
void outPutAuthorityFiles(vector<BigNumber>& fileSet){
	//���n��BigNumber�ļ�	
	for (int i = 0; i < fileSet.size(); i++){
		string SharesPath = "AuthorityShare\\Authority";
		SharesPath = SharesPath + to_string(i);	
		ofstream outfile(SharesPath, ios::binary);
		if (!outfile){
			cerr << "open error!" << endl;
			abort();
		}
		unsigned char zz[BitLen_136 / 8];
		NTL::BytesFromZZ(zz, fileSet[i], BitLen_136 / 8);
		outfile.write((char*)& zz, (BitLen_136 / 8));
		//�ر�����ļ�
		outfile.clear();
		outfile.close();
	}
}
void readShareFiles(vector<File>& shareInFile){
	for (int i = 0; i < shareInFile.size(); i++){
		string SharesPath = "Share\\Sharing";
		SharesPath = SharesPath + to_string(i);
		shareInFile[i].write_share_file_into_word_set(SharesPath);
	}
}

void secretReconstruction_V(ShamirSSScheme& ReAuthorityTool, UInt m, UInt t, BigNumber prime_Au, vector<BigNumber>& AuthorityShare_Set, vector<UInt>& y_id){
	ReAuthorityTool.write_m_nN(m);
	ReAuthorityTool.write_m_nK(t);
	ReAuthorityTool.write_m_bnPrimeNr(prime_Au);
	ReAuthorityTool.write_m_vSharingPart(AuthorityShare_Set);
	ReAuthorityTool.write_x_id_set(y_id);
	ReAuthorityTool.secret_Reconstruction();
}
void genAuShares(vector<BigNumber>& AuthorityShare_Set, BigNumber v, UInt t, UInt m, vector<UInt>& y_id, BigNumber& prime_Au, vector<double>&time_v_share){

	ShamirSSScheme auSS;												//����һ��Shamir's secret sharing�Ĳ���������
	clock_t start, end;
	auSS.set_Parameter(t, m, v, y_id, prime_Au, true);
	start = clock();
	auSS.gen_SecretSharePart();											//����shares
	end = clock();
	outPutTime(start, end, "����share of v", time_v_share);
	AuthorityShare_Set = auSS.get_m_vSharing();							//��ȡvi shares
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
	//����length��word_set���ռ䣻�����뱻�ֳ�N��С������

	BigNumber rePrime;
	unsigned char zz[BitLen_136 / 8];
	cout << "bitlen136:" << (BitLen_136 / 8) << endl;
	infile.read((char*)& zz, (BitLen_136 / 8));		//read size ���ֽ�Ϊ��λ���룻
	NTL::ZZFromBytes(rePrime, zz, BitLen_136 / 8);

	cout << "reprime:" << rePrime << endl;
	//cout << "reprime length:" << NTL::NumBits(rePrime) << endl;
	//cout << "reprime sign:" << NTL::sign(rePrime) << endl;
	return rePrime;

}

void outputResult(ofstream & shuchu, string message, vector<double> & set,Var& varia){

	vector<double> lin;
	sort(set.begin(), set.end());
	for (int i = varia.DelCount; i < set.size() - varia.DelCount; i++){
		lin.push_back(set[i]);
	}
	double count = 0;
	for (int j = 0; j < lin.size(); j++){
		count = count + lin[j];
	}
	shuchu << message << count / lin.size() << endl;
}
void outputResult(ofstream & shuchu, ofstream & sabun, string message, vector<double> & set, Var& varia){

	vector<double> lin;
	sort(set.begin(), set.end());
	for (int i = varia.DelCount; i < set.size() - varia.DelCount; i++){
		lin.push_back(set[i]);
	}
	double count = 0;
	for (int j = 0; j < lin.size(); j++){
		sabun << message << "-" << j << ":" << lin[j] << endl;
		count = count + lin[j];
	}
	shuchu << message << count / lin.size() << endl;
}
void readAuthrotiy(vector<BigNumber>& FileAuthShare_Set, Var& varia){

	FileAuthShare_Set.resize(varia.ThresholdProviders);
	//vector<BigNumber> AuSet(varia.TotalProviders);
	
	for (int i = 0; i < varia.ThresholdProviders; i++){
		string filePath="AuthorityShare\\Authority";
		filePath = filePath + to_string(i);
		ifstream infile(filePath, ios::binary);
		if (!infile){
			cerr << "open error!" << endl;
			abort();
		}
		unsigned char zz[BitLen_136 / 8];
		infile.read((char*)& zz, (BitLen_136 / 8));		//read size ���ֽ�Ϊ��λ���룻
		NTL::ZZFromBytes(FileAuthShare_Set[i], zz, BitLen_136 / 8);
		infile.clear();
		infile.close();

	}

}
void outPutKey(BigNumber key, Var& varia){

	ofstream outfile(varia.KeyPath, ios::binary);
	if (!outfile){
		cerr << "open error!" << endl;
		abort();
	}

	BigNumber rePrime;
	unsigned char zz[BIT256 / 8];
	NTL::BytesFromZZ(zz, key, BIT256 / 8);
	outfile.write((char*)& zz, (BIT256 / 8));		//read size ���ֽ�Ϊ��λ���룻

}
BigNumber readKey(Var& varia){
	ifstream infile(varia.KeyPath, ios::binary);
	if (!infile){
		cerr << "open error!" << endl;
		abort();
	}
	BigNumber rePrime;
	unsigned char zz[BIT256 / 8];
	infile.read((char*)& zz, (BIT256 / 8));		//read size ���ֽ�Ϊ��λ���룻
	NTL::ZZFromBytes(rePrime, zz, BIT256 / 8);
	return rePrime;
}
