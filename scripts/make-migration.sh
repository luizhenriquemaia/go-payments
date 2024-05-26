helpFunction()
{
   echo ""
   echo "Usage: $0 -n complementary name of file"
   echo -e "\t-n complementary name of file"
   exit 1 # Exit script after printing help
}

while getopts ":n:?:" opt
do
    case "$opt" in
        n ) file_name=$(echo "$OPTARG" | sed -e 's/ /_/g');;
        ? ) helpFunction ;;
    esac
done

if [ -z "$file_name" ] || [ "$file_name" == " " ];
then
    echo -e "\nError: file name is required";
    helpFunction
fi

goose -dir "configs/database/migrations" create $file_name sql
exit 1