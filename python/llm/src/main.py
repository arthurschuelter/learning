from src.chat import LlamaChat
import logging
from langchain.prompts import ChatPromptTemplate, PromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough
from langchain.retrievers.multi_query import MultiQueryRetriever

def create_retriever(vector_db, llm):
    """Create a multi-query retriever."""
    QUERY_PROMPT = PromptTemplate(
        input_variables=["question"],
        template="""
            You are an AI language model assistant. Your task is to generate five
            different versions of the given user question to retrieve relevant documents from
            a vector database. By generating multiple perspectives on the user question, your
            goal is to help the user overcome some of the limitations of the distance-based
            similarity search. Provide these alternative questions separated by newlines.
            Original question: {question}""",
    )

    retriever = MultiQueryRetriever.from_llm(
        vector_db.as_retriever(), llm, prompt=QUERY_PROMPT
    )
    logging.info("Retriever created.")
    return retriever

def create_chain(retriever, llm):
    """Create the chain with preserved syntax."""
    # RAG prompt
    template = """
        Answer the question based ONLY on the following context:
        {context}
        Question: {question}
        """

    prompt = ChatPromptTemplate.from_template(template)

    chain = (
        {"context": retriever, "question": RunnablePassthrough()}
        | prompt
        | llm
        | StrOutputParser()
    )

    logging.info("Chain created with preserved syntax.")
    return chain

def main():
    print("RAG Assistant")

    # Initialize the language model
    chat = LlamaChat()
    retriever = create_retriever(chat.vector_db, chat.model)
    chain = create_chain(retriever, chat.model)

    if chat.vector_db is None:
        print("Failed to load or create the vector database.")
        return

    while True:
        user_input = input("You: ")
        if user_input.lower() in ["exit", "quit"]:
            print("Exiting. Goodbye!")
            break
        if user_input:
            try:
                response = chain.invoke(input=user_input)

                print("Chat:", response)
            except Exception as e:
                print(f"An error occurred: {str(e)}")
    
if __name__ == "__main__":
  main()
