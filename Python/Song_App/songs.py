if __name__ == '__main__':

    print("1 : see all songs")
    print("2 : Add a new song")
    print("3 : delete a song")
    print("4 : edit a song")
    print("5 : exit\n")

    songs = ["marry on a cross" , "believer" , "snap" , "way down we go" , "husn"]

    def add_song():
        a = input("please enter a song name which you want to add : ")
        songs.append(a)
        all_songs()

    def delete_song():
        all_songs()
        inp = int(input("please enter a song index which you want to delete : "))
        songs.remove(songs[inp-1])
        all_songs()

    def all_songs():
        for i in range(1,len(songs)+1):
            print(i , ":" , songs[i-1])

    def edit_song():
        all_songs()
        inp = int(input("which song you want to edit please enter an index of that song : "))
        new = input("enter a new name for the song :")
        songs[inp-1] = new
        all_songs()

    inp = int(input("please select any one option : "))
    if (inp == 1):
        all_songs()
    
    if (inp == 2):
        add_song()

    if (inp == 3):
        delete_song()

    if (inp == 4):
        edit_song()

    if (inp == 5):
        print("")