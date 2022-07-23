package main

func main() {
	// conn, err := grpc.Dial("localhost:4500",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }
	// client := pb.NewStudentServiceClient(conn)

	// student := pb.CreateStudentRequest{
	// 	FullName:  "Khafid Prayoga",
	// 	Address:   "Mojokerto",
	// 	Class:     "X-TKJ-2",
	// 	BirthDate: uint64(time.Now().Unix()),
	// }
	// created, err := client.CreateStudent(context.Background(), &student)
	// if err != nil {
	// 	fmt.Printf("error %v", err)
	// }
	// fmt.Printf("Data student\n\n%v \n", created)
	// data, _ := client.GetStudentById(context.Background(), &pb.GetStudentByIdRequest{Id: 1})
	// fmt.Printf("Data student\n\n%v \n", data)

	// data, _ := client.GetAllStudent(context.Background(), &emptypb.Empty{})
	// fmt.Printf("Data student\n\n%v \n", data)

	// data, _ := client.UpdateStudentAddress(context.Background(), &pb.UpdateStudentAddressRequest{
	// 	Id:         1,
	// 	NewAddress: "Germany",
	// })
	// fmt.Printf("Status perubahan data \n\n%v \n", data)

	// data, _ := client.DeleteStudent(context.Background(), &pb.DeleteStudentRequest{Id: 20})
	// fmt.Printf("Status perubahan data \n\n%v \n", data)
}
