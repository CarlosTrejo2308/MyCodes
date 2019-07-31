void main() {
  var auto = new Vehiculo("Ford", "Focus");
  auto.color = "Rojo";
  
  auto.Arrancar();

}

class Vehiculo {
  String color;
  String modelo;
  String marca;
  
  Vehiculo(this.marca, this.modelo);
  
  void Arrancar() {
    print("Hola, soy el auto $marca y estoy acelerando");
  }
  
  void Acelerar(int velocidad) {
    print("Mi velocidad actual es $velocidad km/h");
  }
  
  void CambiarValor(Vehiculo vehiculo) {
    vehiculo.marca = "Hazda";
  }
}
