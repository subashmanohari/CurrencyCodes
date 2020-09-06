import React, { Component } from "react";
import logo from "./logo.svg";
import Table from "react-bootstrap/Table";
import "./App.css";
import { AgGridReact } from "ag-grid-react";
import "ag-grid-community/dist/styles/ag-grid.css";
import "ag-grid-community/dist/styles/ag-theme-alpine.css";
import axios from "axios";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cad: "",
      usd: "",
      hkd: "",
      inr: "",
      btc: "",
      jpy: "",
      kwd: "",
      byn: "",
      pkr: "",
      irr: "",
    };
  }



  interval = null;

  componentDidMount() {
    this.interval = setInterval(this.getData, 60000);
    this.getData();
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  getData = () => {
    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/cad",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ cad: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });

    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/inr",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ inr: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });

    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/btc",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ btc: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });

    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/usd",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ usd: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });

    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/jpy",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ jpy: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });
    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/kwd",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ kwd: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });
    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/byn",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ byn: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });
    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/pkr",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ pkr: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });

    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/irr",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        console.log(response);
        this.setState({ irr: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });
    axios
      .get(
        "https://xe7ulkv238.execute-api.ap-south-1.amazonaws.com/Development/fetchdata/hkd",
        {
          responseType: "json",
        }
      )
      .then((response) => {
        this.setState({ hkd: response.data });
        //  console.log(Object.keys(this.state.tableData.rates));

        // console.log(this.state.tableData.rates.AED);
      });
  }

  render() {
    let tbl = {
      columnDefs: [
        { headerName: "Currency", field: "make" },
        { headerName: "Value (w.r.t EUR)", field: "model", sortable: true },
      ],

      rowData: [
        { make: "CAD", model: this.state.cad },
        { make: "HKD", model: this.state.hkd },
        { make: "BTC", model: this.state.btc },
        { make: "USD", model: this.state.usd },
        { make: "JPY", model: this.state.jpy },
        { make: "KWD", model: this.state.kwd },
        { make: "BYN", model: this.state.byn },
        { make: "PKR", model: this.state.pkr },
        { make: "IRR", model: this.state.irr },
        { make: "INR", model: this.state.inr },
      ],
    };
    //  let tm = JSON.parse(this.state.tableData.rates);
    console.log(this.state);
    return (
      <div
        className="ag-theme-alpine"
        style={{ height: "500px", width: "600px"}}
      >
        <AgGridReact
          columnDefs={tbl.columnDefs}
          rowData={tbl.rowData}
        ></AgGridReact>
      </div>
    );
  }

  //   const petList = Object.entries(this.state.tableData).map(([key, value]) => {
  //     return (
  //       <div>
  //         {key} : {value.toString()}
  //       </div>
  //     );
  //   });
  //   return <div>{petList}</div>;
  // }
}
export default App;