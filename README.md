# TextProbe

TextProbe is a powerful and scalable text search tool designed for efficiently querying large collections of documents. This tool is specifically tailored for searching through Wikipedia abstracts, making it an ideal solution for knowledge exploration and information retrieval.

## Features

- **Efficient Indexing:** TextProbe utilizes advanced indexing mechanisms, allowing it to efficiently handle large datasets. This feature makes it suitable for datasets with up to 1 million documents.

- **Full-Text Search:** Perform full-text searches on Wikipedia abstracts, enabling quick retrieval of relevant information based on your queries.

## Getting Started

Follow these steps to get started with TextProbe:

### Prerequisites

- Make sure you have Go installed on your system.

## Scalability
TextProbe is designed with scalability in mind. Its indexing and search algorithms are optimized to handle large volumes of documents. The tool has been tested with datasets containing up to 1 million documents, ensuring it can scale to meet the demands of extensive knowledge repositories.


### Installation

1. **Clone the Repository:**
   ```bash
   git clone <repository_url>
   cd TextProbe

## Usage

### Download Wikipedia Dump:
Get the latest Wikipedia abstract dump from [https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz) and save it to your preferred location.

### Run TextProbe:
Use the following command to run TextProbe, specifying the dump path and your search query:

```bash
go run main.go -p <path_to_dump> -q <your_search_query>
