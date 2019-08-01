void main() {
  var list = [1, 6, 7, 32, 42];
 
  
  var len = list.length;
  
  for(int i = 0; i < len; i++) {
    print("Indice: $i  Valor: " + list[i].toString());
  }
 
  
  var onlyString = new List<String>();
  
  onlyString.add("Flutter");
  onlyString.add("Dart");
  onlyString.add("es genial");

  len = onlyString.length;
  
  print("");
  
  for(int i = 0; i < len; i++) {
    print("Indice: $i  Valor: " + onlyString[i].toString());
  }
  
  var listPersonas = new List<Persona>();
  
  var personal = new Persona();
  personal.nombre = "Rodrigo";
  
  var persona2 = new Persona();
  persona2.nombre = "Ivan";
  
  listPersonas.add(personal);
  listPersonas.add(persona2);
  
  for(var x in listPersonas) {
    print(x.nombre); 
  }
}

class Persona {
  String nombre;
}
