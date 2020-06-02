package main
import (
    "github.com/tadvi/winc"
    "strconv"
)

var nbCoups int = 0
var tableauMorpion = [9]int{0,0,0,0,0,0,0,0,0}


func main() {
    mainWindow := winc.NewForm(nil)
    mainWindow.SetSize(700, 500)
    mainWindow.SetText("Morpion")
    
    DisplayButton(mainWindow)
    DisplayInformations("Le joueur " + strconv.FormatInt(int64(nbCoups % 2 + 1), 10) + " doit jouer.", mainWindow)
    
    mainWindow.Center()
    mainWindow.Show()
    mainWindow.OnClose().Bind(wndOnClose)
    winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
    winc.Exit()
}

func CreateButton(posX int, posY int, window *winc.Form) {
    btn := winc.NewPushButton(window)
    btn.SetPos(posX, posY)
    btn.SetSize(60, 60)
    btn.SetText("")
    btn.OnClick().Bind(func(e *winc.Event){
        ButtonClicked(btn, window)
    })
}

func DisplayButton(window *winc.Form){
    CreateButton(100,100,window)
    CreateButton(170,100,window)
    CreateButton(240,100,window)
    CreateButton(100,170,window)
    CreateButton(170,170,window)
    CreateButton(240,170,window)
    CreateButton(100,240,window)
    CreateButton(170,240,window)
    CreateButton(240,240,window)
}

func ButtonClicked(button *winc.PushButton, window *winc.Form) {
    if (nbCoups >= 9){
        DisplayInformations("Tous les coups ont été joué, pas de gagnants", window)
    } else {
        if (nbCoups % 2 == 0) {
            button.SetText("X")
            tableauMorpion[nbCoups] = 1
        } else {
            button.SetText("O")
            tableauMorpion[nbCoups] = 2
        }
        VerifGagner(window)
        nbCoups++
    }

}

func DisplayInformations(text string,window *winc.Form){
    textInformations :=winc.NewLabel(window)
    textInformations.SetPos(5,5)
    textInformations.SetText(text)
    textInformations.SetSize(300,30)
}

func VerifGagner(window *winc.Form){
    if (tableauMorpion[0] == 1 && tableauMorpion[1] == 1 && tableauMorpion[2] == 1 ||
        tableauMorpion[3] == 1 && tableauMorpion[4] == 1 && tableauMorpion[5] == 1 ||
        tableauMorpion[6] == 1 && tableauMorpion[7] == 1 && tableauMorpion[8] == 1 ||     
        tableauMorpion[0] == 1 && tableauMorpion[3] == 1 && tableauMorpion[6] == 1 ||
        tableauMorpion[1] == 1 && tableauMorpion[4] == 1 && tableauMorpion[7] == 1 ||
        tableauMorpion[2] == 1 && tableauMorpion[5] == 1 && tableauMorpion[8] == 1 ||      
        tableauMorpion[0] == 1 && tableauMorpion[4] == 1 && tableauMorpion[8] == 1 ||
        tableauMorpion[2] == 1 && tableauMorpion[4] == 1 && tableauMorpion[6] == 1) {
        DisplayInformations("Le joueur 1 à gagner",window)
      } else if (tableauMorpion[0] == 2 && tableauMorpion[1] == 2 && tableauMorpion[2] == 2 ||
        tableauMorpion[3] == 2 && tableauMorpion[4] == 2 && tableauMorpion[5] == 2 ||
        tableauMorpion[6] == 2 && tableauMorpion[7] == 2 && tableauMorpion[8] == 2 || 
        tableauMorpion[0] == 2 && tableauMorpion[3] == 2 && tableauMorpion[6] == 2 ||
        tableauMorpion[1] == 2 && tableauMorpion[4] == 2 && tableauMorpion[7] == 2 ||
        tableauMorpion[2] == 2 && tableauMorpion[5] == 2 && tableauMorpion[8] == 2 ||   
        tableauMorpion[0] == 2 && tableauMorpion[4] == 2 && tableauMorpion[8] == 2 ||
        tableauMorpion[2] == 2 && tableauMorpion[4] == 2 && tableauMorpion[6] == 2) {
        DisplayInformations("Le joueur 2 à gagner",window)
      } else {
        DisplayInformations("Le joueur " + strconv.FormatInt(int64(nbCoups % 2 + 1), 10) + " doit jouer.",window)
        DisplayInformations(strconv.FormatInt(int64(tableauMorpion[0]), 10) + strconv.FormatInt(int64(tableauMorpion[1]), 10) +strconv.FormatInt(int64(tableauMorpion[2]), 10) , window)
      }
}



