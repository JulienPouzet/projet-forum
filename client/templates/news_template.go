package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
)

type NewsTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
	News        []models.News
}

func News(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)

	if err != nil {
		if err == http.ErrNoCookie {
			data := NewsTemplate{
				CurrentUser: models.User{},
				LoggedIn:    false,
				News:        getDummyNewsItems(),
			}

			renderNewsTemplate(w, "web/pages/news.html", data)
			return
		}

		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		log.Printf("Error fetching user: %v", err)
		return
	}

	data := NewsTemplate{
		CurrentUser: currentUser,
		LoggedIn:    true,
		News:        getDummyNewsItems(),
	}

	renderNewsTemplate(w, "web/pages/news.html", data)
}

// UTILS FUNCTIONS

func renderNewsTemplate(w http.ResponseWriter, templateFile string, data NewsTemplate) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}
}

func getDummyNewsItems() []models.News {
	return []models.News{
		{UUID: "1", Title: "Hit me hard and soft", Content: "Billie Eilish, the global pop sensation, has once again captivated her fans with her latest album titled 'Hit Me Hard and Soft'. This bold and emotional opus confirms the undeniable talent of the young artist. A blend of emotions and depth 'Hit Me Hard and Soft' immerses the listener in an introspective journey through the ups and downs of life. The poignant lyrics and Billie Eilish's enchanting voice take the listener on a whirlwind of emotions, oscillating between strength and vulnerability. Innovative musical production The musical production of this album is as remarkable as Billie Eilish's vocal performances. The subtle arrangements and unique sounds create an immersive atmosphere that perfectly complements the artist's profound lyrics. In conclusion, 'Hit Me Hard and Soft' is an album that resonates with the audience, offering a captivating and immersive musical experience. Billie Eilish continues to push the boundaries of creativity and emotion through this opus that is sure to leave a lasting impression on the minds and hearts of her fans worldwide.", Date: "2024-05-28", Category: "Pop", Author: "Louna"},
		{UUID: "2", Title: "Radical Optimism", Content: "Dua Lipa, the international pop sensation, has enchanted her fans with the release of her latest album titled 'Radical Optimism'. This bold and energetic opus testifies to the talent and creativity of the British singer. A message of positivity and hope 'Radical Optimism' is an album that exudes joy and positivity. The songs on the album convey a message of hope and optimism, inviting the listener to see the bright side of life and embrace each moment with optimism. Catchy and modern sounds The music of 'Radical Optimism' is a blend of catchy pop sounds and modern beats that encourage dancing. Dua Lipa explores new musical horizons while staying true to her unique style, creating a dynamic and captivating album. In conclusion, 'Radical Optimism' is an album that radiates good vibes and celebrates life and music in all their splendor. Dua Lipa continues to charm her audience with catchy tracks and positive messages, confirming her place among the most influential artists in the current music industry.", Date: "2024-05-04", Category: "Pop", Author: "Maxence"},
		{UUID: "3", Title: "Cowboy Carter", Content: "Beyonce, the pop music icon, once again dazzled the world with the release of her latest album titled 'Cowboy Carter'. This bold and innovative opus testifies to the singer's constant artistic evolution. An exploration of new musical horizons 'Cowboy Carter' marks a turning point in Beyoncé's career by exploring country and folk sounds that the pop diva had not yet touched. The tracks on the album skillfully blend pop and country influences, creating a unique and captivating musical universe. Deep and engaging lyrics The lyrics of 'Cowboy Carter' are both poetic and engaging, addressing universal themes such as love, freedom, and emancipation. Beyoncé delivers emotional vocal performances that transport the listener to the heart of the stories told in each song. In conclusion, 'Cowboy Carter' is an artistic work that testifies to Beyoncé's creative genius and her ability to constantly reinvent herself. This album offers a musical experience rich in emotions and authenticity, once again confirming Beyoncé's place among the greatest artists of her generation.", Date: "2024-03-29", Category: "Country", Author: "Jeanne"},
		{UUID: "4", Title: "Blue lips", Content: "Schoolboy Q, the talented American rapper, recently released his sixth album titled 'Blue Lips'. This album stands out for its authenticity and eclecticism, offering listeners a unique and captivating musical experience. A raw and daring album 'Blue Lips' aims to be an even more raw and daring album than Schoolboy Q's previous projects. With sharp lyrics and innovative sounds, the artist pushes the boundaries of rap and presents a musical universe rich in emotions. Critical reception Critics praise Schoolboy Q's boldness and commitment in this album. The deep lyrics and refined production of 'Blue Lips' have managed to captivate rap and urban music fans, once again confirming the artist's talent for storytelling through his music. In conclusion, 'Blue Lips' is an album that showcases Schoolboy Q's talent and creativity. With this album, the artist continues to make a mark in the music industry with his unique style and ability to innovate, offering listeners an unforgettable musical experience.", Date: "2024-03-01", Category: "Rap", Author: "Melvin"},
		{UUID: "5", Title: "American Dream", Content: "The American rapper 21 Savage shook the music industry with the release of his acclaimed album 'American Dream'. This powerful and introspective opus immerses the listener in the raw realities of life in the United States, offering an unfiltered look at the challenges and triumphs of the contemporary American experience. Exploration of Social Themes 'American Dream' tackles the social themes that shape today's American society in a frank and direct manner. 21 Savage delivers impactful lyrics that address street violence, social injustice, the struggle for survival, and the pursuit of success despite obstacles. Artistic Collaboration The album highlights 21 Savage's collaboration with other renowned artists in the music industry, adding additional dimensions to his compelling narrative. Featuring artists from diverse backgrounds enriches the album's sonic diversity and enhances its emotional impact. Critical and Public Reception 'American Dream' has been praised by critics for its sincerity, urban poetry, and ability to capture the essence of modern American experience. The album has also achieved commercial success, attracting a wide audience and solidifying 21 Savage's position as a prominent artist of his generation. In conclusion, 'American Dream' by 21 Savage is more than just a rap album; it is a powerful testimony to life, struggles, and dreams in the United States. Through his music, 21 Savage provides an authentic voice to the often overlooked realities of American society, making this album a must-listen for fans of conscious and socially engaged rap.", Date: "2024-01-12", Category: "Rap", Author: "Angélique"},
	}
}
