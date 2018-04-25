package lib

import (
	"fmt";
	"os";
	"log";
	"io/ioutil";
)


func FileOperations() {

	fmt.Println("\n\n ## The following content is exploring Go File Operations ## \n");
	
	file, err := os.Create("sampleFile.txt");

	if err != nil {
		log.Fatal(err);
	}

	file.WriteString("This is some random text");
	file.Close();

	stream, err := ioutil.ReadFile("sampleFile.txt");

	if err != nil {
		log.Fatal(err);
	}

	readString := string(stream);

	fmt.Println(readString);

}