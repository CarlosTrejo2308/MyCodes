import 'package:flutter/material.dart';


class MyButton extends StatefulWidget {

  @override
  _MyButtonState createState() => _MyButtonState();
}


class _MyButtonState extends State<MyButton> {
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Statful Widget")
      )
    ,);
  }
}
