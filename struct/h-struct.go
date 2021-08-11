package main

import "fmt"

// buatlah struct song
// dengan title, artist string
type song struct {
	title  string
	artist string
}

// buatlah playlist struct
// dengan genre string, songTitles []string
// songArtist []string, dan songs []song
type playlist struct {
	genre       string
	songTitlesb []string
	songArtist  []string
	songs       []song
}

func main() {

	// Buat dua variabel song1, dan song2 menggunakan struct song
	song1 := song{ title:  "Aku ingin Pulang", artist: "Ebiet G. Ade"}
	song2 := song{ title:  "Rembulan menangis", artist: "Ebiet G. Ade"}
	// misalkan artist BTS dengan 2 lagu yang berbeda

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// copy song2 kedalam song1
	song1 = song2

	// gunakan if untuk membandingkan title dan artis
	// jika sama print songs are equal
	// kalau tidak print songs are not equal
	if song1.artist == song2.artist && song1.title == song2.title {
		fmt.Println("songs are equal")
	} else {
		fmt.Println("songs are not equal")
	}

	// buatlah slice of songs
	slceOfSongs := []song{
		song1,
		{"sajkd", "asdh"},
	}
	// buatlah playlist misalkan rock dan masukkan
	// slice of songs tersebut ke dalam playlist rock
	rock := playlist{
		genre: "rock",
	songTitlesb: []string{"A", "B"},
	songArtist:  []string{"C", "D"},
	songs: slceOfSongs,
	}

	// iterasi songs di dalam rock kemudian print
	for _, val := range rock.songs {
		fmt.Println(val)
	}

}
