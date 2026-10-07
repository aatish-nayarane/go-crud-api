package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"vm-crud/internal/model"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

func (c *Client) CreateVM(req model.CreateVMRequest) (*model.VM, error) {
	url := fmt.Sprintf("%s/vms", c.baseURL)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	res, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	res.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(res)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	var vm model.VM
	if err := json.NewDecoder(response.Body).Decode(&vm); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &vm, nil
}

func (c *Client) ListVMs() ([]model.VM, error) {
	url := fmt.Sprintf("%s/vms", c.baseURL)
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	client := &http.Client{}
	response, err := client.Do(res)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	var vms []model.VM
	if err := json.NewDecoder(response.Body).Decode(&vms); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return vms, nil
}

func (c *Client) GetVMByID(id int) (*model.VM, error) {
	url := fmt.Sprintf("%s/vms/%d", c.baseURL, id)
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	client := &http.Client{}
	response, err := client.Do(res)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	var vm model.VM
	if err := json.NewDecoder(response.Body).Decode(&vm); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &vm, nil
}

func (c *Client) UpdateVM(id int, req model.UpdateVMRequest) (*model.VM, error) {
	url := fmt.Sprintf("%s/vms/%d", c.baseURL, id)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	res, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	res.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(res)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()
	fmt.Println("----------Response Status-------:", response.Status)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var vm model.VM
	if err := json.NewDecoder(response.Body).Decode(&vm); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &vm, nil
}

func (c *Client) DeleteVM(id int) error {
	url := fmt.Sprintf("%s/vms/%d", c.baseURL, id)
	res, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	client := &http.Client{}
	response, err := client.Do(res)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	return nil
}


// Does the value need to survive the current stack frame?
//                  │
//         ┌────────┴────────┐
//        NO                 YES
//        ↓                   ↓
//      STACK               HEAP
// we will  learn  about escape  analysis  in  go  programming  language  and  how  to  use  it  to  improve  the  security  of  our  applications.
func createVMA() *model.CreateVMRequest {
	//
	data := model.CreateVMRequest{
		Name: "Test VM",
		CPU:  2,
		RAM:  4096,
	}
	return &data

}

// that will  be stored in  heap memory becouse  of  we taking  pointer or returning  pointer  from  function  and  that  will  be  stored in  heap memory and  that  will  be  used in  main function
func createVMB() model.CreateVMRequest {
	//
	data := model.CreateVMRequest{
		Name: "Test VM",
		CPU:  2,
		RAM:  4096,
	}
	return data

}

// that will  be stored in  stack memory becouse  of  we returning  value from  function and  that  will  be  stored in  stack memory and  that  will  be  used in  main function
func createVMc(data *model.CreateVMRequest) {
	//
	data.CPU = 4

}

type test struct {
	Name string
	Age  int
}


func CreateTest(t interface{}) {
	fmt.Printf("Type of t: %T\n", t)
	if testStruct, ok := t.(test); ok {
		fmt.Printf("Name: %s, Age: %d\n", testStruct.Name, testStruct.Age)
	} else {
		fmt.Println("t is not of type test")
	}
}
func processVM(vm *model.VM) {
    fmt.Println(vm.Name)
}

func main() {
	// Request model call

	// client := NewClient("http://localhost:8080/v1")
	req := model.CreateVMRequest{
		Name: "Test VM",
		CPU:  2,
		RAM:  4096,
	}
	// vm, err := client.CreateVM(req)
	// if err != nil {
	// 	fmt.Printf("Error creating VM: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Created VM: %+v\n", vm)
	// //  list  vms  here
	// vms, err := client.ListVMs()
	// if err != nil {
	// 	fmt.Printf("Error listing VMs: %v\n", err)
	// 	return
	// }
	// fmt.Printf("List of VMs: %+v\n", vms)
	// //  get  vm  by  id  here
	// id :=  int(vm.ID)
	// vmByID, err := client.GetVMByID(id)
	// if err != nil {
	// 	fmt.Printf("Error getting VM by ID: %v\n", err)
	// 	return
	// }
	// fmt.Printf("VM by ID: %+v\n", vmByID)
	// //  update  vm  here
	// updateReq := model.UpdateVMRequest{
	// 	Name: "Updated Test VM",
	// 	CPU:  4,
	// 	RAM:  8192,
	// }
	// updatedVM, err := client.UpdateVM(id, updateReq)
	// if err != nil {
	// 	fmt.Printf("Error updating VM: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Updated VM: %+v\n", updatedVM)
	// //  delete  vm  here
	// err = client.DeleteVM(id)
	// if err != nil {
	// 	fmt.Printf("Error deleting VM: %v\n", err)
	// 	return
	// }
	// fmt.Println("Deleted VM successfully")
	// //  list  vms  again  to  confirm  deletion
	// vms, err = client.ListVMs()
	// if err != nil {
	// 	fmt.Printf("Error listing VMs after deletion: %v\n", err)
	// 	return
	// }
	// fmt.Printf("List of VMs after deletion: %+v\n", vms)

	// esacape  analysis data
	createVMAData := createVMA()
	fmt.Printf("CreateVMA Data: %+v\n", createVMAData)

	createVMc(&req)
	fmt.Printf("CreateVMA Data after modification: %+v\n", createVMAData)

	// var sata test
	sata := test{
		Name: "John Doe",
		Age:  30,
	}
	CreateTest(sata)

	// process  vm  pointer  here
	vm := &model.VM{
		Name: "Test VM",
		CPU:  2,
		RAM:  4096,
	}
	processVM(vm)

}
