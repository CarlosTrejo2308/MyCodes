import 'package:flutter/material.dart';

void main() {
  var app = MaterialApp(
    home: Scaffold(
      appBar: AppBar(
        title: Text("Mi primera applicacion"),
      ),
      body: Container(
        child: Center(
          child: Text("Hello, World!", 
          style: TextStyle(fontSize: 40.0),
          ),
        ),
      ),

      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        onPressed: () {print("Click");},
      ),
    ),
  );

  runApp(app);

}
