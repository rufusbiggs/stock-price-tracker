# Real-Time Stock Price Tracker and ML Recommendation System  

## Project Overview  
This project focuses on building a **Real-Time Stock Price Tracker** as a foundation for a future **Machine Learning Recommendation System**. The ultimate goal is to provide a robust system for tracking, analyzing, and making data-driven predictions about stock prices.  

---

## Project Roadmap  

### Phase 1: Real-Time Stock Price Tracker  
The initial phase focuses on building the data pipeline and ensuring scalability. This phase includes:  

#### Key Features:  
- **Stock Price Fetcher**:  
  - Scheduled AWS Lambda functions to fetch real-time stock data from APIs like Alpha Vantage or IEX Cloud.  
  - Historical stock price data stored in a PostgreSQL database.  

- **API Layer**:  
  - REST APIs for:  
    - Fetching historical stock data.  
    - Fetching the latest stock price.  

- **Data Streaming**:  
  - Redis Streams to provide real-time price updates to subscribers.  

- **Data Cleanup Automation**:  
  - Scheduled Lambda functions for optimizing the database and cleaning old data.  

- **Monitoring and CI/CD**:  
  - AWS CloudWatch for monitoring.  
  - AWS CodePipeline for automated deployments.  

#### Why Phase 1 Matters:  
- Builds a scalable and well-structured pipeline for fetching, processing, and storing stock data.  
- Collects historical stock price data essential for training ML models in the next phase.  
- Provides hands-on experience with **AWS tools** and **Golang** development.  

---

### Phase 2: Simple ML Recommendation System  
Building on the tracker, this phase introduces a basic **ML-based buy/sell/hold recommendation system**.  

#### Steps:  

1. **Define Features for Recommendations**:  
   - Calculate technical indicators from historical data:  
     - **SMA** (Simple Moving Average).  
     - **EMA** (Exponential Moving Average).  
     - **RSI** (Relative Strength Index).  
     - **MACD** (Moving Average Convergence Divergence).  

2. **Rule-Based Recommendation System**:  
   - Implement a baseline rules-based system:  
     - Example rules:  
       - RSI < 30 → "Buy".  
       - RSI > 70 → "Sell".  
       - MACD crosses above the signal line → "Buy".  
       - SMA(50) > SMA(200) → "Buy".  

3. **Transition to ML**:  
   - Collect and label historical stock data stored in PostgreSQL.  
   - Train a basic supervised learning model using **scikit-learn** or **TensorFlow/Keras** to classify buy/sell/hold signals.  

---

## Technology Stack  
- **Backend**: Golang for API and data processing.  
- **Database**: PostgreSQL for storing historical data.  
- **Real-Time Streaming**: Redis Streams for live updates.  
- **Cloud Services**:  
  - AWS Lambda for scheduled tasks and cleanup.  
  - AWS CloudWatch for monitoring.  
  - AWS CodePipeline for CI/CD.  
- **Machine Learning** (Phase 2): Python, Pandas, scikit-learn, TensorFlow/Keras.  

---

## Getting Started  
1. Clone the repository:  
   ```bash
   git clone https://github.com/your-username/real-time-stock-tracker.git
   cd real-time-stock-tracker
   ```
2. Set up environment variables for AWS credentials and API keys.
3. Deploy the backend using AWS CodePipeline or a local environment.  

---

## Future Work  
- Improve ML models for higher accuracy using advanced techniques like LSTMs or Reinforcement Learning.  
- Integrate sentiment analysis using news or social media data.  
- Add user authentication and customizable stock alerts.  

---

## Contributions  
Contributions are welcome! Feel free to open issues or submit pull requests.  
