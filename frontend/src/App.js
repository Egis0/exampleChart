import React, { Component } from 'react';
import './App.css';
import axios from 'axios';
import moment from 'moment'
import { Line } from 'react-chartjs-2';

const colors = ['rgb(155, 202, 200)', 'rgb(176, 195, 222)', 'rgb(123, 248, 222)'];
const colorsBg = ['rgba(175, 222, 220, .5)', 'rgba(196, 215, 242, .5)', 'rgba(145, 168, 164, .5)'];

const articleStyle = {
  padding: '0 120px'
};

const headingStyle = {
  marginBottom: 10,
  marginTop: 35,
  fontWeight: 600,
  fontStyle: 'italic'
};

class App extends Component {
  state = { corrections: [] };

  componentDidMount() {
    axios.get('http://localhost:8080/corrections')
        .then(response => this.setState({ corrections: response.data }));
  }

  getValues(label) {
    return this.state.corrections
        .filter(correction => correction.title === label)
        .map(correction => correction.value);
  }

  render() {
    const labels = [...new Set(this.state.corrections.map(correction => moment(correction.date).format('ddd, DD MMM YYYY')))];
    const datasetLabels = [...new Set(this.state.corrections.map(correction => correction.title))];

    const datasets = datasetLabels.map((label, index) => {
      return {
        label,
        data: this.getValues(label),
        borderColor: colors[index],
        backgroundColor: colorsBg[index]
      };
    });

    const options = {
      scales: {
        xAxes: [{
          stacked: true
        }],
        yAxes: [{
          stacked: true
        }]
      },
      legend: {
        position: 'bottom',
        labels: {
          boxHeight: 15,
          boxWidth: 15
        }
      }
    };

    const data = {
      labels,
      datasets
    };

    return (
      <div className="App">
        <h4 style={headingStyle}>Correction statistics test</h4>
        <article className="canvas-container" style={articleStyle}>
          <Line data={data} options={options}/>
        </article>
      </div>
    );
  }
}

export default App;
