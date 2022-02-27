class EmployeeStruct {
    private name: string;
    private id: number;

    public getName(): string {
        return this.name;
    }

    public setName(name: string): void {
        this.name = name;
    }

    public getId(): number {
        return this.id;
    }

    public setId(id: number): void {
        this.id = id;
    }

}

const employee = new EmployeeStruct();
employee.setId(1);
employee.setName('Name');

console.log(employee.getId());
console.log(employee.getName());