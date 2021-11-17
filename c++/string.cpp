#include <iostream>
#include <string>

using namespace std;

int main()
{
    string s1 = "Hello";
    string s2 = "World";
    string s3 = s1 + " " + s2;

    char text[] = "Hello World";

    cout << s3 << endl;
    cout << s3.size() << endl;

    cout << text << endl;

    string texto = "10";
    int numero = stoi(texto);
    cout << numero << endl;
    return 0;
}