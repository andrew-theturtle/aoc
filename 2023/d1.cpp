#include <stdio.h>
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int trebuchet1(string line) {
    string res;
    for (char c : line) {
        if (isdigit(c)) {
            res.push_back(c);
            break;
        }
    }

    for (int i = line.size() - 1; i >= 0; i--) {
        if (isdigit(line[i])) {
            res.push_back(line[i]);
            break;
        }
    }

    if (!res.size()) return 0;

    return stoi(res);
}

int isNamedInt(string x) {
    if (x.find("one") == 0) return 1;
    if (x.find("two") == 0) return 2;
    if (x.find("three") == 0) return 3;
    if (x.find("four") == 0) return 4;
    if (x.find("five") == 0) return 5;
    if (x.find("six") == 0) return 6;
    if (x.find("seven") == 0) return 7;
    if (x.find("eight") == 0) return 8;
    if (x.find("nine") == 0) return 9;
    return -1;
}

pair<int, int> findFirstNamedInt(string line) {
    for (int i = 0; i < line.size(); i++) {
        string x = line.substr(i, min(5, (int) line.size() - i));
        int res = isNamedInt(x);
        if (res != -1) return make_pair(i, res); 
    }
    return make_pair(-1, 0);
}

pair<int, int> findLastNamedInt(string line) {
    for (int i = line.size(); i >= 0; i--) {
        string x = line.substr(i, min(5, (int) line.size() - i));
        int res = isNamedInt(x);
        if (res != -1) return make_pair(i, res);
    }
    return make_pair(-1, 0);
}

pair<int, int> findFirstInt(string line) {
    for (int i = 0; i < line.size(); i++) {
        if (isdigit(line[i])) {
            return make_pair(i, line[i] - '0');
        }
    }
    return make_pair(-1, 0);
}

pair<int, int> findLastInt(string line) {
    for (int i = line.size() - 1; i >= 0; i--) {
        if (isdigit(line[i])) {
            return make_pair(i, line[i] - '0');
        }
    }
    return make_pair(-1, 0);
}

int trebuchet2(string line) {
    if (!line.size()) return 0;
    auto fNamedInt = findFirstNamedInt(line);
    auto lNamedInt = findLastNamedInt(line);
    auto fInt = findFirstInt(line);
    auto lInt = findLastInt(line);

    int firstInt, lastInt;
    if (fNamedInt.first == -1 || fInt.first == -1) {
        if (fNamedInt.first == -1) {
            firstInt = fInt.second;
        } else {
            firstInt = fNamedInt.second;
        }
    } else {
        if (fNamedInt.first < fInt.first) {
            firstInt = fNamedInt.second;
        } else {
            firstInt = fInt.second;
        }
    }

    if (lNamedInt.first == -1 || lInt.first == -1) {
        if (lNamedInt.first == -1) {
            lastInt = lInt.second;
        } else {
            lastInt = lNamedInt.second;
        }
    } else {
        if (lNamedInt.first > lInt.first) {
            lastInt = lNamedInt.second;
        } else {
            lastInt = lInt.second;
        }
    }

    return firstInt * 10 + lastInt;
}

int main() {
    ifstream file("inputs/d1.txt");
    int res = 0;
    if (file.is_open()) {
        string line;
        while (getline(file, line)) {
            res += trebuchet2(line);
        }
    }

    cout << res << endl;

    return 0;
}

