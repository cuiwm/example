package main

import "fmt"

// inline uint32 address_hash(const char* addr)        {
//     uint32 seed = 131;
//     uint32 hash = 0;
//     const char* p = addr;
//     while(*p){
//         hash = hash * seed + (*p++);
//              }
//     return hash % 0x7FFFFFFF;
// }

// bool CycleTableDB::GetTableNameByAddr(const std::string& addr, const string &originalName, std::string &tablename){    if (originalName.empty()) return false;

//     uint32 id = Utility::address_hash(addr.c_str());

//     stringstream ss;    ss << originalName << "_" << (id % 1000);    tablename = ss.str();    return true;}

func address_hash(addr string) uint32 {
	var (
		seed uint32 = 131
		hash uint32 = 0
	)

	for i := 0; i < len(addr); i++ {
		hash = hash*seed + uint32(addr[i])
	}

	// const char* p = addr;
	// while(*p){
	//     hash = hash * seed + (*p++);
	// 		}

	return hash % 0x7FFFFFFF
}

func main() {

	//originalName  << "_" << (id % 1000);
	tablename := fmt.Sprintf("%s_%d", "UserIncome", address_hash("serkjb"))
	//tablename = ss.str();    return true;}
	fmt.Println("Hello", tablename)

	copy()
}
