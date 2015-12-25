{- Welcome to your first Elm program

Read up on syntax:
  http://elm-lang.org/learn/Syntax.elm

Learn about the Elm's core libraries:
  http://package.elm-lang.org/packages/elm-lang/core/latest/

-}

import Graphics.Element exposing (..)
import Debug


main : Element
main =
    flow down
      [ helloWorld
      , welcomeGraphics
      ]

main : Element
main =
  let
    one = Debug.log "one" 1
    two = Debug.log "two" 2
    three = Debug.log "three" 3
  in
    show "Hello"

helloWorld : Element
helloWorld =
    show "Hello, World!"


welcomeGraphics : Element
welcomeGraphics =
    let dimensions = 90
        imgSize = 30
        elmLogo =
          image imgSize imgSize "http://elm-lang.org/logo.png"

        elmsPerSide = dimensions // imgSize

        row =
          flow right (List.repeat elmsPerSide elmLogo)
    in
        flow down (List.repeat elmsPerSide row)
