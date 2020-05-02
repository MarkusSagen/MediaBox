import React from 'react';
import logo from './logo.svg';
import './App.css';


class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			name: 			"",
			greeting: 	""
		};
	}

	handleChange = (e) => {
		this.setState({ name: e.target.value })
	}

	handleSubmit = (e) => {
		e.preventDefault();
		fetch(`api?name=${encodeURIComponent(this.state.name)}`)
			.then(resp => resp.json())
			.then(state => this.setState(state));
	}


	render() {
  	return (
    	<div className="App">
      	<header className="App-header">
        	<img src={logo} className="App-logo" alt="logo" />
				</header>
				<main>
					<div>
						<form onSubmit={this.handleSubmit}>
							<label htmlFor="name"> Enter name </label>
							<input 
								id="name"
								type="text" 
								value={this.state.name}
								onChange={this.handleChange}
							/>
							<button type="submit">Submit</button>
						</form>
						<p>{this.state.greeting}</p>
					</div>
				</main>
   		</div>
  	);
	}
}

export default App;
