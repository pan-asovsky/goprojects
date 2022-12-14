package main

import "fmt"

func main() {

	fmt.Print("1 для ручного расчёта, 2 для автоматического: ")
	var input = InputInt()

	if input == 1 {
		ReductionManualCalculate()
	} else if input == 2 {
		ReductionAutoCalculate()
	} else {
		fmt.Print("Неверный ввод!")
	}

}

func ReductionAutoCalculate() {

	fmt.Print("Расчёт до n единиц брони: ")
	var armor = InputInt()

	fmt.Print("Введите модификатор(0.00 - 1.00): ")
	var mdf = InputFloat32()

	for i := 1; i <= armor; i++ {
		fmt.Println("Броня:", IntToString(i)+". Cнижение:", CalculateDamageReduction(float32(i), mdf)+"%")
	}

}

func ReductionManualCalculate() {

	fmt.Print("Количество единиц брони: ")
	var armor = InputFloat32()

	fmt.Print("Введите модификатор(0.00 - 1.00): ")

	var percent = InputFloat32()
	var reduction = CalculateDamageReduction(armor, percent)

	fmt.Println("Снижение урона:", reduction+"%")
}

func CalculateDamageReduction(armor, percent float32) string {

	var reductionPercent = (armor * percent) / (1 + armor*percent)
	var reductionInt = int(reductionPercent * 100)
	return IntToString(reductionInt)
}
