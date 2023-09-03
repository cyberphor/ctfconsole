// import logo from './logo.svg';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Navigation from './components/Navigation';
import Footer from './components/Footer'
import ButtonGetPlayers from './components/Buttons';

export default function App() {
  return (
    <div className="App">
      <Navigation/>
      <ButtonGetPlayers/>
      <Footer/>
    </div>
  );
}
