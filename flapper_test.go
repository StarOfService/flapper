package flapper

import (
  // "fmt"
  "reflect"
  "testing"
  // "time"
)

func TestMarshal(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
    e string
    F []string
    G [3]string
    H []int
    I [3]int
    K []bool
  }

  var suite = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
    e: "non-public fields should be skipped",
    F: []string{"aa", "bb", "cc"},
    G: [3]string{"aa", "bb", "cc"},
    H: []int{23, 54, 76},
    I: [3]int{23, 54, 76},
    K: []bool{true, false, false},
  }

  assert := map[string]string{
    "A": "a-value",
    "B": "2",
    "C": "true",
    "D.DA": "d-value",
    "D.DB": "3.14E+00",
    "F.0": "aa",
    "F.1": "bb",
    "F.2": "cc",
    "G.0": "aa",
    "G.1": "bb",
    "G.2": "cc",
    "H.0": "23",
    "H.1": "54",
    "H.2": "76",
    "I.0": "23",
    "I.1": "54",
    "I.2": "76",
    "K.0": "true",
    "K.1": "false",
    "K.2": "false",
  }

  serial, err := Marshal(suite)
  if err != nil {
    t.Errorf(err.Error())
  }

  for k, _ := range assert {
    if assert[k] != serial[k] {
      t.Errorf("Expected %v value for key %v, but got %v", assert[k], k, serial[k])
    }
  }
}


func TestMarshalDelimiterPrefix(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
    e string
  }

  var suite = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
    e: "non-public fields should be skipped",
  }

  assert := map[string]string{
    "test:A": "a-value",
    "test:B": "2",
    "test:C": "true",
    "test:D:DA": "d-value",
    "test:D:DB": "3.14E+00",
  }

  fh, err := New("test", ":")
  if err != nil {
    t.Errorf(err.Error())
  }
  serial, err := fh.Marshal(suite)
  if err != nil {
    t.Errorf(err.Error())
  }

  for k, _ := range assert {
    if assert[k] != serial[k] {
      t.Errorf("Expected %v value for key %v, but got %v", assert[k], k, serial[k])
    }
  }
}

// func TestMarshalEmbedded(t *testing.T) {

//   type TStruct2 struct {
//     CA string
//     CB int32
//   }
  
//   type TStruct1 struct {
//     TStruct2
//     A string
//     B int
//   }

//   suite := TStruct1{
//     TStruct2{"ca-value", 214748364},
//     "a-value",
//     2,
//   }

//   assert := map[string]string{
//     "A": "a-value",
//     "B": "2",
//     "CA": "ca-value",
//     "CB": "214748364",
//   }

//   serial, err := Marshal(suite)
//   if err != nil {
//     t.Errorf(err.Error())
//   }

//   for k, _ := range assert {
//     if assert[k] != serial[k] {
//       t.Errorf("Expected %v value for key %v, but got %v", assert[k], k, serial[k])
//     }
//   }
// }


func TestUnmarshal(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
    F []string
    G [3]string
    H []int
    I [3]int
    K []bool
  }

  suite := map[string]string{
    "A": "a-value",
    "B": "2",
    "C": "true",
    "D.DA": "d-value",
    "D.DB": "3.14E+00",
    "F.0": "aa",
    "F.1": "bb",
    "F.2": "cc",
    "G.0": "aa",
    "G.1": "bb",
    "G.2": "cc",
    "H.0": "23",
    "H.1": "54",
    "H.2": "76",
    "I.0": "23",
    "I.1": "54",
    "I.2": "76",
    "K.0": "true",
    "K.1": "false",
    "K.2": "false",
  }

  var assert = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
    F: []string{"aa", "bb", "cc"},
    G: [3]string{"aa", "bb", "cc"},
    H: []int{23, 54, 76},
    I: [3]int{23, 54, 76},
    K: []bool{true, false, false},
  }

  var deserial TStruct1
  err := Unmarshal(suite, &deserial)
  if err != nil {
    t.Errorf(err.Error())
  }

  if !reflect.DeepEqual(assert, deserial) {
    t.Errorf("Test suite object %v doesn't match to the generated data %v", assert, deserial)
  }
}

func TestUnmarshalDelimiterPrefix(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
  }

  suite := map[string]string{
    "test:A": "a-value",
    "test:B": "2",
    "test:C": "true",
    "test:D:DA": "d-value",
    "test:D:DB": "3.14E+00",
  }

  var assert = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
  }

  fh, err := New("test", ":")
  if err != nil {
    t.Errorf(err.Error())
  }

  var deserial TStruct1
  err = fh.Unmarshal(suite, &deserial)
  if err != nil {
    t.Errorf(err.Error())
  }

  if !reflect.DeepEqual(assert, deserial) {
    t.Errorf("Test suite object %v doesn't match to the generated data %v", assert, deserial)
  }
}

func TestMarshalUnmarshal(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
  }

  var suite = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
  }

  serial, err := Marshal(suite)
  if err != nil {
    t.Errorf(err.Error())
  }

  var deserial TStruct1
  err = Unmarshal(serial, &deserial)
  if err != nil {
    t.Errorf(err.Error())
  }

  if !reflect.DeepEqual(suite, deserial) {
    t.Errorf("Test suite object %v doesn't match to the generated data %v", suite, deserial)
  }
}

func TestMarshalUnmarshalDelimiterPrefix(t *testing.T) {

  type TStruct2 struct {
    DA string
    DB float32
  }
  
  type TStruct1 struct {
    A string
    B int
    C bool
    D TStruct2
  }

  var suite = TStruct1{
    A: "a-value",
    B: 2,
    C: true,
    D: TStruct2{
      DA: "d-value",
      DB: 3.14,
    },
  }

  fh, err := New("test", ":")
  if err != nil {
    t.Errorf(err.Error())
  }

  serial, err := fh.Marshal(suite)
  if err != nil {
    t.Errorf(err.Error())
  }

  var deserial TStruct1
  err = fh.Unmarshal(serial, &deserial)
  if err != nil {
    t.Errorf(err.Error())
  }

  if !reflect.DeepEqual(suite, deserial) {
    t.Errorf("Test suite object %v doesn't match to the generated data %v", suite, deserial)
  }
}