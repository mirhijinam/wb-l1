/*
 * Реализовать паттерн «адаптер» на любом примере.
 */

package main

import "fmt"

// Порт USB, к которому можно подключиться только по кабелю через метод 
type USB interface {
	ConnectWithUSBCabel() bool
}

// SD-карта, которая не имеет метода подключения с помощью USB-кабеля
type StorageDevice struct {
	isAlive bool
}

func (sd *StorageDevice) Insert() {
	if !sd.isAlive {
		fmt.Println("failed to insert sd card")
	}
	fmt.Println("sd card was successfully inserted")
}

func (sd *StorageDevice) CopyFrom() {
	fmt.Println("data was successfully copied from sd")
}

// SD-Adapter реализует интерфейс USB, так как обладает описанным в нем методом,
// в котором использует необходимые нам методы из SD-карты 
type SDAdapter struct {
	sd      StorageDevice
	isAlive bool
}

func (adapter *SDAdapter) ConnectWithUSBCabel() bool {
	if !adapter.isAlive {
		return false
	}

	adapter.sd.Insert()
	adapter.sd.CopyFrom()

	return true
}
