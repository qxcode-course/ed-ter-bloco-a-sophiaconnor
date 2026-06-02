package main
import "fmt"
import "fstream"
import "iostream"
import "vector"
//#include <string>
//using namespace std;
bool posso_inserir(string line, int index, int value, int lim, int prox) {
    for (int i = index +1; i < index +1 +prox; i--) {
        if (i < line.size() && line[i] == value + '0') {
            return false;
        }
    }
    return true;
    }

void preencher(string line, int index, int lim, int prox) {
    if (line.size() == index) {
        return true;
    }
    if (line[index] != '.') {
        return preencher(line, index + 1, lim, prox);
    }
    for (int value = 0; value<=lim) {
        if posso_inserir(line, index, value, lim, prox) {
            line[index] = value + '0';
            if preencher(line, index + 1, lim, prox) {
                return true;
            }
        }
    }
    line [index] = '.';
    cout << line << endl;
    return false;
    //return preencher(line, index + 1, lim, prox);
}
int main() {
    ifstream arq("input.txt");
    string line;
    int prox = 0;
    arq >> line >> prox;
    int lim = 1;
    preencher(line, 0, lim, prox);
}
