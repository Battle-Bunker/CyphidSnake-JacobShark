package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
  _ "github.com/BattlesnakeOfficial/rules"
)

func HeuristicSnakeAntiCramming(snapshot agent.GameSnapshot) float64 {
  // Step 1. Get the value of the corners
  chunk_1_open_space := 0
  chunk_2_open_space := 0
  chunk_3_open_space := 0
  chunk_4_open_space := 0
  
  snake_head_in_chunk_1 := snapshot.You().Head().X < 5 && snapshot.You().Head().Y < 5
  snake_head_in_chunk_2 := snapshot.You().Head().X > 4 && snapshot.You().Head().Y < 5
  snake_head_in_chunk_3 := snapshot.You().Head().X < 5 && snapshot.You().Head().Y > 4
  snake_head_in_chunk_4 := snapshot.You().Head().X > 4 && snapshot.You().Head().Y > 4

  if snake_head_in_chunk_1 {
    for _, cellColumn := range snapshot.Board().Cells {
      for _, cell := range cellColumn {
        if cell.Coordinates().X < 5 && cell.Coordinates().Y < 5 {
          if cell.Kind() == agent.CellEmpty {
            chunk_1_open_space++
          }
        }
      }
    }
  }


  if snake_head_in_chunk_2 {
    for _, cellColumn := range snapshot.Board().Cells {
      for _, cell := range cellColumn {
        if cell.Coordinates().X > 4 && cell.Coordinates().Y < 5 {
          if cell.Kind() == agent.CellEmpty {
            chunk_2_open_space++
          }
        }
      }
    }
  }


  if snake_head_in_chunk_3 {
    for _, cellColumn := range snapshot.Board().Cells {
      for _, cell := range cellColumn {
        if cell.Coordinates().X < 5 && cell.Coordinates().Y > 4 {
          if cell.Kind() == agent.CellEmpty {
            chunk_3_open_space++
          }
        }
      }
    }
  }

  if snake_head_in_chunk_4 {
    for _, cellColumn := range snapshot.Board().Cells {
      for _, cell := range cellColumn {
        if cell.Coordinates().X > 4 && cell.Coordinates().Y > 4 {
          if cell.Kind() == agent.CellEmpty {
            chunk_4_open_space++
          }
        }
      }
    }
  }



  // if snake_head_in_chunk_1 {
  //   return float64(chunk_1_open_space)
  // }

  // if snake_head_in_chunk_2 {
  //   return float64(chunk_2_open_space)
  // }

  // if snake_head_in_chunk_3 {
  //   return float64(chunk_3_open_space)
  // }


  // if snake_head_in_chunk_4 {
  //   return float64(chunk_4_open_space)
  // }

  // Step 2.1 Get the value of the corners
  corner_1_value := float64(chunk_1_open_space)
  corner_2_value := float64(chunk_2_open_space)
  corner_3_value := float64(chunk_3_open_space)
  corner_4_value := float64(chunk_4_open_space)

  // Step 2.2 Do the math to get the fractional parts for our sqaure

  // var totals []int

  // try using this:
  // var corner_fraction_Y [4]float64  
  // if you use this in a loop, you should be able to condense this code, ask an AI if it feels too complicated
  // totals = append(totals, int(corner_1_value))

  corner_1_fraction_Y := corner_1_value / (float64(snapshot.You().Head().Y) / float64(snapshot.Height()))
  corner_1_fraction_X := corner_1_value / (float64(snapshot.You().Head().X) / float64(snapshot.Width()))
  corner_1_total := corner_1_fraction_Y + corner_1_fraction_X


  corner_2_fraction_Y := corner_2_value / float64(snapshot.You().Head().Y/snapshot.Height())
  corner_2_fraction_X := corner_2_value / float64(snapshot.You().Head().X/snapshot.Width())
  corner_2_total := corner_2_fraction_Y + corner_2_fraction_X
  
  corner_3_fraction_Y := corner_3_value / float64(snapshot.You().Head().Y/snapshot.Height())
  corner_3_fraction_X := corner_3_value / float64(snapshot.You().Head().X/snapshot.Width())
  corner_3_total := corner_3_fraction_Y + corner_3_fraction_X

  corner_4_fraction_Y := corner_4_value / float64(snapshot.You().Head().Y/snapshot.Height())
  corner_4_fraction_X := corner_4_value / float64(snapshot.You().Head().X/snapshot.Width())
  corner_4_total := corner_4_fraction_Y + corner_4_fraction_X

  // var corners []float64 = []float64{corner_1_value, corner_2_value, corner_3_value, corner_4_value}

  // loop_counter
  

  total := (corner_1_total + corner_2_total + corner_3_total + corner_4_total) / 4 // average the results
  if total == 0 {
    // add some logging here idk
  }
  

  
  return float64(total)
}
