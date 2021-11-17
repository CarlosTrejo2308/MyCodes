#include <stdio.h>

#define PIZZA_COST 1.5
const char NEWLINE = '\n';

int main() {
    int pizza_count;
    double total_cost;

    printf("How many pizzas do you want to order? ");
    scanf("%d", &pizza_count);

    total_cost = pizza_count * PIZZA_COST;

    printf("Your total cost is $%.2f%c", total_cost, NEWLINE);

    return 0;
}
