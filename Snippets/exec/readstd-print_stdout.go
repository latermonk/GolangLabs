 

// READ
stdout, err := cmd.StdoutPipe()
  if err != nil {
  log.Fatal(err)
  }

//PRINT
for {
		fmt.Println("==========")
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
