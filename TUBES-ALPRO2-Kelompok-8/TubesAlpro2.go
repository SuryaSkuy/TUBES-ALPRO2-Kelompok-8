package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --- Global Variables ---
// Menggunakan array statis (bukan slice) untuk menyimpan data utama seperti pengguna, teman, status, dan komentar.
// Variabel global hanya digunakan untuk array utama seperti `users`, yang memuat semua data pengguna.
const MAX_USERS = 100
const MAX_FRIENDS = 50
const MAX_STATUSES = 50
const MAX_COMMENTS = 50

type User struct {
	Username    string
	Password    string
	Profile     string
	Friends     [MAX_FRIENDS]string
	Statuses    [MAX_STATUSES]string
	Comments    [MAX_STATUSES][MAX_COMMENTS]string
	FriendCount int
	StatusCount int
}

var users [MAX_USERS]User
var userCount int
var loggedInIndex = -1

// Utility Functions
// `findUserIndex` menggunakan algoritma pencarian sequential untuk mencari pengguna berdasarkan username.
// Fungsi ini modular dan digunakan di beberapa bagian aplikasi.
func findUserIndex(username string) int {
	// Loop iterasi linear untuk mencari indeks pengguna berdasarkan username.
	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			return i
		}
	}
	return -1
}

func login(username, password string) int {
	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			return i
		}
	}
	return -1
}

// Sorting Functions
// `selectionSortUsersByUsername` menggunakan algoritma selection sort untuk mengurutkan data pengguna.
// Mendukung urutan ascending dan descending sesuai parameter.
func selectionSortUsersByUsername(order string) {
	for i := 0; i < userCount-1; i++ {
		idx := i
		for j := i + 1; j < userCount; j++ {
			// Menggunakan kondisi untuk menentukan urutan berdasarkan parameter `order`.
			if (order == "asc" && users[j].Username < users[idx].Username) || (order == "desc" && users[j].Username > users[idx].Username) {
				idx = j
			}
		}
		if idx != i {
			// Swap elemen jika diperlukan.
			users[i], users[idx] = users[idx], users[i]
		}
	}
}

// Core Functions
// --- Modular Subprograms ---
// Setiap fitur diimplementasikan dalam fungsi modular dengan parameter yang jelas.
// Contoh: Fungsi `register` menangani proses registrasi pengguna baru.
func register() {
	if userCount >= MAX_USERS {
		fmt.Println("User limit reached.")
		return
	}

	// Input username dan password.
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	// Validasi apakah username sudah ada dengan memanfaatkan `findUserIndex`.
	if findUserIndex(username) != -1 {
		fmt.Println("Username already exists.")
		return
	}

	// Menambahkan pengguna baru ke array `users`.
	users[userCount] = User{Username: username, Password: password}
	userCount++
	fmt.Println("Registration successful!")
}

func loginHandler() {
	if loggedInIndex != -1 {
		fmt.Println("Already logged in.")
		return
	}

	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	index := login(username, password)
	if index == -1 {
		fmt.Println("Invalid username or password.")
		return
	}

	loggedInIndex = index
	fmt.Printf("Welcome, %s!\n", users[loggedInIndex].Username)
}

func logout() {
	if loggedInIndex == -1 {
		fmt.Println("No user is logged in.")
		return
	}

	fmt.Printf("Goodbye, %s!\n", users[loggedInIndex].Username)
	loggedInIndex = -1
}

func updateProfile() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	fmt.Print("Enter new profile information: ")
	reader := bufio.NewReader(os.Stdin)
	profile, _ := reader.ReadString('\n')
	users[loggedInIndex].Profile = strings.TrimSpace(profile)
	fmt.Println("Profile updated!")
}

func addFriend() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	if users[loggedInIndex].FriendCount >= MAX_FRIENDS {
		fmt.Println("Friend list is full.")
		return
	}

	var friendUsername string
	fmt.Print("Enter the username of the friend to add: ")
	fmt.Scanln(&friendUsername)

	friendIndex := findUserIndex(friendUsername)
	if friendIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	if friendUsername == users[loggedInIndex].Username {
		fmt.Println("You cannot add yourself as a friend.")
		return
	}

	for i := 0; i < users[loggedInIndex].FriendCount; i++ {
		if users[loggedInIndex].Friends[i] == friendUsername {
			fmt.Println("This user is already your friend.")
			return
		}
	}

	users[loggedInIndex].Friends[users[loggedInIndex].FriendCount] = friendUsername
	users[loggedInIndex].FriendCount++
	fmt.Println("Friend added successfully!")
}

func postStatus() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	if users[loggedInIndex].StatusCount >= MAX_STATUSES {
		fmt.Println("Status list is full.")
		return
	}

	fmt.Print("Enter your status: ")
	reader := bufio.NewReader(os.Stdin)
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)
	users[loggedInIndex].Statuses[users[loggedInIndex].StatusCount] = status
	users[loggedInIndex].StatusCount++
	fmt.Println("Status posted!")
}

func commentOnStatus() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	var username string
	fmt.Print("Enter the username of the user whose status you want to comment on: ")
	fmt.Scanln(&username)

	userIndex := findUserIndex(username)
	if userIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	if users[userIndex].StatusCount == 0 {
		fmt.Println("This user has no statuses.")
		return
	}

	fmt.Println("--- User's Statuses ---")
	for i := 0; i < users[userIndex].StatusCount; i++ {
		fmt.Printf("%d. %s\n", i+1, users[userIndex].Statuses[i])
	}

	var statusIndex int
	fmt.Print("Select a status number to comment on: ")
	fmt.Scanln(&statusIndex)

	if statusIndex < 1 || statusIndex > users[userIndex].StatusCount {
		fmt.Println("Invalid status number.")
		return
	}

	fmt.Print("Enter your comment: ")
	reader := bufio.NewReader(os.Stdin)
	comment, _ := reader.ReadString('\n')
	comment = strings.TrimSpace(comment)

	statusIndex-- // Convert to zero-based index
	comments := &users[userIndex].Comments[statusIndex]
	for i := 0; i < MAX_COMMENTS; i++ {
		if (*comments)[i] == "" {
			(*comments)[i] = fmt.Sprintf("%s: %s", users[loggedInIndex].Username, comment)
			fmt.Println("Comment added!")
			return
		}
	}

	fmt.Println("Comments for this status are full.")
}

func viewHome() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	fmt.Println("--- Home ---")
	for i := 0; i < userCount; i++ {
		fmt.Printf("%s:\n", users[i].Username)
		for j := 0; j < users[i].StatusCount; j++ {
			fmt.Printf("  - %s\n", users[i].Statuses[j])
			for k := 0; k < MAX_COMMENTS; k++ {
				if users[i].Comments[j][k] != "" {
					fmt.Printf("    * %s\n", users[i].Comments[j][k])
				}
			}
		}
	}
}

// --- Edit, Search, and Delete Operations ---
// Menggunakan algoritma sequential search untuk mencari pengguna atau mengedit data mereka.
// Contoh: `searchUser` memungkinkan pengguna mencari profil berdasarkan username.
func searchUser() {
	if loggedInIndex == -1 {
		fmt.Println("Please login first.")
		return
	}

	// Input username untuk pencarian.
	var username string
	fmt.Print("Enter username to search: ")
	fmt.Scanln(&username)

	// Menggunakan `findUserIndex` untuk menemukan data.
	userIndex := findUserIndex(username)
	if userIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	// Menampilkan data pengguna yang ditemukan.
	fmt.Printf("--- User Profile ---\nUsername: %s\nProfile: %s\n", users[userIndex].Username, users[userIndex].Profile)
	fmt.Println("Friends:")
	for i := 0; i < users[userIndex].FriendCount; i++ {
		fmt.Printf("  - %s\n", users[userIndex].Friends[i])
	}
}

// --- Sorting Implementation ---
// Selection sort sudah diimplementasikan di `selectionSortUsersByUsername`.
// Untuk mendukung kebutuhan lain, algoritma insertion sort dapat ditambahkan.

// --- Menu ---
// Mengintegrasikan semua fitur melalui menu utama yang modular.
// Setiap pilihan menu memanggil subprogram terkait.
func menu() {
	for {
		fmt.Println("\n--- Social Media App ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Logout")
		fmt.Println("4. Update Profile")
		fmt.Println("5. Add Friend")
		fmt.Println("6. Post Status")
		fmt.Println("7. View Home")
		fmt.Println("8. Comment on Status")
		fmt.Println("9. Search User")
		fmt.Println("10. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		// Memanggil subprogram sesuai pilihan.
		switch choice {
		case 1:
			register()
		case 2:
			loginHandler()
		case 3:
			logout()
		case 4:
			updateProfile()
		case 5:
			addFriend()
		case 6:
			postStatus()
		case 7:
			viewHome()
		case 8:
			commentOnStatus()
		case 9:
			searchUser()
		case 10:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// --- Main Function ---
// Fungsi utama hanya memanggil menu utama.
func main() {
	menu()
}
