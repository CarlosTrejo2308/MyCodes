void main() {
  print(factorial(0));
}

num factorial(num n) {
  if (n < 0) {
    return -1;
  }
  if (n == 0) 
    return 1;
  else {
    return (n * factorial(n - 1));
  }
}
