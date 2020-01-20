provider "music" {
  spotify_user = "124907240"
}

data "music_concert" "recently_played_artists" {
  max_concerts = 2
}

variable "email" {
  type = string
  default = "pedrampejman2010@gmail.com"
}

resource "calendar_event" "concert" {
  count       = length(data.music_concert.recently_played_artists.concerts)
  summary     = data.music_concert.recently_played_artists.concerts[count.index].title
  description = data.music_concert.recently_played_artists.concerts[count.index].title
  location    = data.music_concert.recently_played_artists.concerts[count.index].location
  start       = data.music_concert.recently_played_artists.concerts[count.index].date_start
  end         = data.music_concert.recently_played_artists.concerts[count.index].date_end

  attendee {
    email = var.email
  }
}
