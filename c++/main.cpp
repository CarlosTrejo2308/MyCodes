#include <iostream>

using namespace std;

int main()
{
    int a, b, c;
    cin >> a >> b >> c;
    if (a == b && b == c)
        cout << "EQUILATERAL";
    else if (a == b || b == c || a == c)
        cout << "ISOSCELES";
    else
        cout << "SCALENE";

    const char *str = "Hello World";
    cout << str << endl;

    int lista_edades[] = {10, 20, 30, 40, 50};

    // loop for each lista edades and print it
    for (int i = 0; i < 5; i++)
    {
        cout << lista_edades[i] << endl;
    }


    return 0;
}