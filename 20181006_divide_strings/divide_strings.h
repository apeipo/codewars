#ifndef INC_20181006_DIVIDE_STRINGS_DIVIDE_STRINGS_H
#define INC_20181006_DIVIDE_STRINGS_DIVIDE_STRINGS_H

#include <iostream>
#include <vector>
#include <string>

using namespace std;

string sub_string(string a, string b) {
    string r = "";
    int fix = 0;
    for (int i = a.size() - 1, j = b.size() - 1; i >= 0; i--, j--) {
        int tmp = (a[i] - fix) - (j < 0 ? '0' : b[j]);
        fix = ((tmp >= 0) ? 0 : 1);
        r = to_string((tmp >= 0) ? tmp : tmp + 10) + r;
    }
    r = r.erase(0, r.find_first_not_of('0'));
    return r.empty() ? "0" : r;
}

vector<string> divide_strings(string a, string b) {
    if (a == "0") {
        return vector<string>{"0", "0"};
    }
    a.erase(0, a.find_first_not_of('0'));
    b.erase(0, b.find_first_not_of('0'));
    if (a.size() < b.size()) {
        return vector<string>{"0", a};
    }
    auto idx = 0;
    string r = "", l = "", part = "";
    while (idx < a.size()) {
        part = (part == "0" ? "" : part);

        //fix part until part.size == b.size
        while (true) {
            part += a[idx++];
            if(part.size() >= b.size() || idx >= a.size()) {
                break;
            }
            r += "0";
        }

        part.erase(0, part.find_first_not_of('0'));
        //do (part - b) until part < b
        uint n = 0; // [0, 9]
        while (part.size() > b.size() || (part.size() == b.size() && part >= b)) {
            part = sub_string(part, b);
            n++;
        }
        r += to_string(n);
        l = part;
    }

    r.erase(0, r.find_first_not_of('0'));
    return vector<string>{r.empty() ? "0" : r, l.empty() ? "0" : l};
}

#endif //INC_20181006_DIVIDE_STRINGS_DIVIDE_STRINGS_H
