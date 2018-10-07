#include <iostream>
#include <vector>
#include <sstream>
#include "divide_strings.h"

using namespace std;

void doTest(string a,   string b, vector<string> expect) {
    auto r = divide_strings(a, b);
    assert(r.size() == 2);

    ostringstream rs;
    rs << r[0] << "," << r[1];
    ostringstream es;
    es << expect[0] << "," << expect[1];

    if (rs.str() != es.str()) {
        cout << "Failed:   " << "expect \n[" << rs.str() << "] to equal \n[" << es.str() << "]" << endl;
    } else {
        cout << a << "/" << b << "=" << es.str() << "     success" << endl;
    }
}

void doSubStringTest(string a, string b, string expect) {
    string r = sub_string(a, b);
//    cout << "res:"<< r << endl;
    if (r != expect) {
        cout << "Failed:   " << "expect [" << r << "] to equal [" << expect << "]" << endl;
    }
}

int main() {
    doTest("0", "5", vector<string>{"0", "0"});
    doTest("4", "5", vector<string>{"0", "4"});
    doTest("10", "2", vector<string>{"5", "0"});
    doTest("219", "11", vector<string>{"19", "10"});
    doTest("23023", "23", vector<string>{"1001", "0"});
    doTest("600001", "100", vector<string>{"6000", "1"});
    doTest("1000", "10", vector<string>{"100", "0"});
    doTest("358435882500682747211879782382309190117565156052767737287194643276725066432592666846227661464548369979884116143020599177390544630285821114848269087358053316636991018456854224677314828252225", "24",
           vector<string>{"14934828437528447800494990932596216254898548168865322386966443469863544434691361118592819227689515415828504839292524965724606026261909213118677878639918888193207959102368926028221451177176", "1"});
//    doSubStringTest("5", "4", "1");
//    doSubStringTest("456", "123", "333");
    doSubStringTest("456", "129", "327");
    doSubStringTest("1456", "1452", "4");
    doSubStringTest("8200824045303269669937", "8100824045303269669937", "100000000000000000000");
}
