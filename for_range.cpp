#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
    int lista[] = {100, 200, 300};
    for (auto &&i : lista)
    {
        cout << i << endl;
    }

    return 0;
}