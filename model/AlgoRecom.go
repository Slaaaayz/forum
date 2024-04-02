package models

import (
	"fmt"
	"math/rand"
	"time"
)

func LoadPost(uid string) []TPost {
	var PostToLoad []TPost
	if GetNbLikesUser(uid) > 30 {
		PostToLoad = append(PostToLoad, GetRandomPost(2)...)
		PostToLoad = append(PostToLoad, GetRecommendPost(8, uid)...)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(PostToLoad), func(i, j int) {
			PostToLoad[i], PostToLoad[j] = PostToLoad[j], PostToLoad[i]
		})
	} else {
		PostToLoad = GetRandomPost(10)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(PostToLoad), func(i, j int) {
			PostToLoad[i], PostToLoad[j] = PostToLoad[j], PostToLoad[i]
		})
	}
	return PostToLoad
}

func GetRandomPost(NbPost int) []TPost {
	var likes = GetMoyLikes() / 3
	println("nb likes apres division :", likes)
	var AllPosts = GetAllPosts()
	if len(AllPosts) == 0 {
		return []TPost{}
	}
	var fini = true
	var RandomPost []TPost
	var NbBoucle int
	rand.Seed(time.Now().UnixNano())
	for fini {
		NbBoucle++

		random := rand.Intn(len(AllPosts))
		println("lepost random : ", random)
		if AllPosts[random].Likes >= likes || NbBoucle > 50 {
			RandomPost = append(RandomPost, AllPosts[random])
		}
		if len(RandomPost) == NbPost {
			fini = false
		}
	}
	return RandomPost
}

func GetMoyLikes() int {
	var likes int
	err := DB.QueryRow("SELECT COALESCE(SUM(likes), 0) FROM post ").Scan(&likes)
	if err != nil {
		panic(err)
	}
	return likes
}

func GetRecommendPost(nbPost int, uid string) []TPost {
	//merci chatgpt pour la requete

	query := `
	SELECT tag, COUNT(*) AS count
	FROM (
		SELECT tags1ID AS tag FROM post JOIN likes ON post.id = likes.idpost WHERE likes.uid = ? AND tags1ID != ''
		UNION ALL
		SELECT tags2ID AS tag FROM post JOIN likes ON post.id = likes.idpost WHERE likes.uid = ? AND tags2ID != ''
		UNION ALL
		SELECT tags3ID AS tag FROM post JOIN likes ON post.id = likes.idpost WHERE likes.uid = ? AND tags3ID != ''
		UNION ALL
		SELECT tags4Id AS tag FROM post JOIN likes ON post.id = likes.idpost WHERE likes.uid = ? AND tags4ID != ''
		UNION ALL
		SELECT tags5ID AS tag FROM post JOIN likes ON post.id = likes.idpost WHERE likes.uid = ? AND tags5ID != ''
	) AS tags
	GROUP BY tag
	`
	//
	rows, _ := DB.Query(query, uid, uid, uid, uid, uid)
	defer rows.Close()
	var TagsLikes = make(map[string]float64)
	var total float64
	var RecommendedPost []TPost
	for rows.Next() {
		var tag string
		var count int
		if err := rows.Scan(&tag, &count); err != nil {
			fmt.Println("Erreur lors de la lecture des résultats:", err)
			return []TPost{}
		}
		total += float64(count)
		TagsLikes[tag] = float64(count)
	}
	// fmt.Println(TagsLikes)
	for tag := range TagsLikes {
		TagsLikes[tag] = (TagsLikes[tag] / total) * 100
	}
	// fmt.Println(TagsLikes)

	for i := 0; i < nbPost; i++ {
		randomNumber := rand.Float64() * 100
		var PostTag string
		var PourcentPlus float64
		for tag, pourcent := range TagsLikes {
			PourcentPlus += pourcent
			if randomNumber < PourcentPlus {
				PostTag = tag
			}
		}
		RecommendedPost = append(RecommendedPost, GetaPostByTag(PostTag, uid))
	}
	return RecommendedPost
}

func GetaPostByTag(tag string, uid string) TPost {
	var TabPost []TPost
	rows, err := DB.Query("SELECT id, name, post, date,IdTopic, image from post WHERE tags1id = ? or tags2id = ? or tags3id = ? or tags4id = ? or tags5id = ?", tag, tag, tag, tag, tag)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var post string
		var date string
		var idtopic int
		var image string
		var Post TPost
		err = rows.Scan(&id, &name, &post, &date, &idtopic, &image)
		if err != nil {
			panic(err)
		}
		Post.Date = date
		Post.Id = id
		Post.IdTopic = idtopic
		Post.Image = image
		Post.IsLiked = IsLiked(uid, id)
		Post.Likes = GetNbLikesPost(id)
		Post.Name = name
		Post.Post = post
		Post.IdUser = GetUser(name).Id
		Post.Pdp = GetUser(name).Image
		TabPost = append(TabPost, Post)
	}
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(TabPost))
	return TabPost[random]
}

func TakeAboPost(uid string) []TPost {
	query := `
        SELECT post.id, post.name, post.post, post.date, post.idtopic, post.image
        FROM post
        JOIN topic ON post.idtopic = topic.id
        JOIN abo ON topic.id = abo.idtopic
        WHERE abo.Uid = ?
    `
	rows, err := DB.Query(query, uid)
	if err != nil {
		fmt.Println(err)
		return []TPost{}
	}
	defer rows.Close()
	var AllPost []TPost
	var GoodPost []TPost

	for rows.Next() {
		var id int
		var name string
		var post string
		var date string
		var idtopic int
		var image string
		var ThePost TPost

		err = rows.Scan(&id, &name, &post, &date, &idtopic, &image)
		if err != nil {
			fmt.Println(err)
			return []TPost{}
		}
		ThePost.Date = date
		ThePost.Id = id
		ThePost.IdTopic = idtopic
		ThePost.IdUser = GetUser(name).Id
		ThePost.Image = image
		ThePost.IsLiked = IsLiked(uid, id)
		ThePost.Likes = GetNbLikesPost(id)
		ThePost.Name = GetUser(name).Pseudo
		ThePost.Pdp = GetUser(name).Image
		ThePost.Post = post
		println(ThePost.Name)
		AllPost = append(AllPost, ThePost)
	}
	if len(AllPost) == 0 {
		return []TPost{}
	}
	for i := 0; i < 10; i++ {
		GoodPost = append(GoodPost, AllPost[rand.Intn(len(AllPost))])
	}
	return GoodPost
}
