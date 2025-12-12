package main

import "fmt"

func main() {
	var bb, tb float64
	var gender string
	var umur int
	var usia bool

	fmt.Println("Kalkulator BMI")
	for {
		for {
			fmt.Print("Jenis Kelamin (L/P): ")
			fmt.Scan(&gender)

			if gender == "L" || gender == "P" {
				break
			} else {
				fmt.Println("Input tidak valid!Mohon masukkan 'L' untuk laki-laki atau 'P' untuk perempuan.")
			}
		}

		for {
			fmt.Print("Apakah usia anda lebih dari atau sama dengan 18 tahun? (true/false): ")
			fmt.Scan(&usia)
			fmt.Print("Berapa usia anda: ")
			fmt.Scan(&umur)

			if umur >= 18 && usia {
				break
			} else if umur > 0 && umur < 18 && !usia {
				fmt.Println("Perhitungan BMI mungkin tidak akurat karena umur di bawah 18 tahun.")
				fmt.Print("Lanjutkan perhitungan BMI? (ya/tidak): ")
				var lanjut string
				fmt.Scan(&lanjut)

				if lanjut == "ya" {
					break
				} else if lanjut == "tidak" {
					fmt.Println("Program selesai, terima kasih.")
					return
				} else {
					fmt.Println("Input tidak valid! Mohon masukkan 'ya' atau 'tidak'.")
				}

			} else {
				fmt.Println("Input tidak valid! Mohon masukkan variabel yang valid!.")
			}
		}

		for {
			fmt.Print("Berat Badan (kg): ")
			fmt.Scan(&bb)

			fmt.Print("Tinggi Badan (cm): ")
			fmt.Scan(&tb)

			if bb > 0 && tb > 0 {
				break
			} else {
				fmt.Println("Input tidak valid, mohon masukkan angka yang valid!")
			}
		}

		bmi := bb / ((tb / 100) * (tb / 100))
		fmt.Printf("BMI Anda: %.2f\n", bmi)

		if bmi < 18.5 {
			fmt.Println("Kategori: Kurus")
		} else if bmi < 25 {
			fmt.Println("Kategori: Normal")
		} else if bmi < 30 {
			fmt.Println("Kategori: Gemuk")
		} else {
			fmt.Println("Kategori: Obesitas")
		}

		var ulang string
		for {
			fmt.Print("Hitung lagi? (ya/tidak): ")
			fmt.Scan(&ulang)
			if ulang == "ya" {
				main()
				return
			} else if ulang == "tidak" {
				fmt.Println("Program selesai. Terima kasih!")
				return
			} else {
				fmt.Println("Input tidak valid! Masukkan hanya ya atau tidak.")
			}
		}

	}
	fmt.Println()
}
