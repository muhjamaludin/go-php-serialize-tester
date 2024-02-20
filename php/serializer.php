<?php

class Serializer
{
    public function serializerNumber(int $number)
    {
        return serialize($number);
    }

    public function unSerializerNumber(string $baseSerialize)
    {
        return unserialize($baseSerialize);
    }

    public function serializerString(string $str)
    {
        return serialize($str);
    }

    public function unSerializerString(string $baseSerialize)
    {
        return unserialize($baseSerialize);
    }

    public function serializeArray(array $arr)
    {
        return serialize($arr);
    }

    public function unSerializeArray(string $baseSerialize)
    {
        return unserialize($baseSerialize);
    }

    public function serializeArrayAssociatif(array $arr)
    {
        return serialize($arr);
    }

    public function unSerializeArrayAssociatif(string $baseSerialize)
    {
        return unserialize($baseSerialize);
    }
}
