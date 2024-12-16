package testdata

const (
	FailMessage = "Invalid domain name format"
)

var SecureURLs = []string{
	"https://google.com",
	"https://youtube.com",
	"https://facebook.com",
	"yandex.ru",
	"https://vk.com",
	"instagram.com",
	"https://wikipedia.org",
	"https://amazon.com",
	"wildberries.ru",
	"https://yahoo.com",
	"https://linkedin.com",
	"reddit.com",
	"mail.ru",
	"twitter.com",
	"avito.ru",
	"https://twitch.tv",
	"aliexpress.ru",
	"https://apple.com",
	"gazeta.ru",
}

var UnsecureURLs = []string{
	"http://google.com",
	"http://httpstat.us",
	"http://neverssl.com",
	"http://httpforever.com",
}

var WrongURLs = []string{
	"htt:/google.com",
	"google",
}
