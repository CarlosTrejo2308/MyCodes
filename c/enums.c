#include <stdio.h>

enum weekDays {
    Monday,
    Tuesday,
    Wednesday,
    Thursday,
    Friday,
    Saturday,
    Sunday
};

enum deck {
    club = 1,
    diamond,
    heart,
    spade
} card;


int main() {
    enum weekDays day;
    day = Sunday;
    printf("Day %d\n", day + 1);

    card = club;
    printf("Card %d\n", card);
    printf("Size of var %ld", sizeof(card));
    return 0;
}