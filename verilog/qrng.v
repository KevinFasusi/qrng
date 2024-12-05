// Count qubits
module qubit_counter (
  //measured qubit input
  input  [3:0] qubit, 

  // out leds add other outputs
  output reg [3:0] led
);

  // nets used to designate two inputs from the qbits from PMOD as RESET and
  // CLOCK set high and low from MCU (arduino nano 33 iot) 
  wire rst0;
  wire rst1;
  wire clk0;
  wire clk1;

  // Reset is the inverse of the pinout from MCU so without digital signal the 
  // input is set to high
  assign rst0 = ~qubit[0];
  assign rts1 = ~qubit[1];
  assign clk0 = ~qubit[2];
  assign clk1 = ~qubit[3];

  // The count for basis state |0> representing the orthonormal state 0 in
  // classical computing
  always @( posedge rst0 or posedge clk0) begin
    if (rst0 == 1'b1) begin
      led <= 4'b0;
    end else begin
      led <= led - 1'b1;
    end 
  end

  // The count for basis state |1> representing the orthonormal state 1 in
  // classical computing
  always @( posedge rst1 or posedge clk1) begin
     if (rst1 == 1'b1) begin
       led <= 4'b0;
     end else begin
       led <= led + 1'b1;
     end
   end
endmodule